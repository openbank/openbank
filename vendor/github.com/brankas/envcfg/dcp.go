package envcfg

import (
	"crypto/tls"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
)

// diskCertProvider provides a certificate provider that watches a local file.
type diskCertProvider struct {
	config *Envcfg

	cert *tls.Certificate

	certPath string
	keyPath  string

	sync.RWMutex
}

// newDiskCertProvider creates a disk cert provider that watches dirPath for
// cert and key.
func newDiskCertProvider(certPath, keyPath string, logf, errf func(string, ...interface{})) (*diskCertProvider, error) {
	dcp := &diskCertProvider{
		certPath: certPath,
		keyPath:  keyPath,
	}
	if err := dcp.loadCertAndKey(); err != nil {
		return nil, err
	}

	// create watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	// don't close, as it runs forever
	// defer watcher.Close()

	addWatch := func(path string) error {
		if err := watcher.Add(certPath); err != nil {
			return err
		}

		// if this is a valid symlink, watch its target too
		if path, err := filepath.EvalSymlinks(certPath); err == nil && path != certPath {
			if err := watcher.Add(path); err != nil {
				return err
			}
		}

		return nil
	}

	go func() {
		for {
			select {
			case <-watcher.Events:
				if err := dcp.loadCertAndKey(); err != nil {
					errf("could not load cert and key: %v", err)
				} else {
					logf("loaded updated certificate")
				}

				// in case we're dealing with symlinks and the target changed,
				// make sure we continue to watch properly
				if err := addWatch(certPath); err != nil {
					errf("could not add watch: %v", err)
				}

			case err := <-watcher.Errors:
				errf("%v", err)
			}
		}
	}()

	// watch the first certificate file
	if err := addWatch(certPath); err != nil {
		return nil, err
	}

	return dcp, nil
}

// loadCertAndKey tries to load the cert and key from disk.
func (dcp *diskCertProvider) loadCertAndKey() error {
	cert, err := tls.LoadX509KeyPair(dcp.certPath, dcp.keyPath)
	if err != nil {
		return err
	}

	dcp.Lock()
	defer dcp.Unlock()

	dcp.cert = &cert
	return nil
}

// GetCertificate satisfies the CertificateProvider interface.
func (dcp *diskCertProvider) GetCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	dcp.RLock()
	defer dcp.RUnlock()

	return dcp.cert, nil
}
