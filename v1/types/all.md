# Annex

##  Address

Address hold all details about Address

| Name        | Type   | Description                                     |
|-------------|--------|-------------------------------------------------|
| CountryName | string | CountryName holds the country name information. |
| CityName    | string | CityName holds the city name information.       |
| State       | string | State holds the state information.              |
| Street      | string | Street holds the street information.            |
| PostalCode  | string | PostalCode holds the postal code information.   |

##  Amount

Amount defines a transaction amount.

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

##  Location

Location is An object representing a latitude/longitude pair. This is expressed as a pairof doubles representing degrees latitude and degrees longitude.Values must be within normalized ranges.Refer to https://godoc.org/google.golang.org/genproto/googleapis/type/latlng.

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

##  OfflineUserInfo

| Name       | Type     | Description                                                |
|------------|----------|------------------------------------------------------------|
| UserID     | string   | UserID                                                     |
| FirstName  | string   | FirstName of the person                                    |
| MiddleName | string   | MiddleName or middle names (space separated) of the person |
| LastName   | string   | LastName or last names (space separated) of the person     |
| MobileNo   | string   | MobileNo contact of the person                             |
| Location   | Location | Location is the physical location of the interaction.      |

##  Profile

Structure of customer profile information

| Name             | Type    | Description                                      |
|------------------|---------|--------------------------------------------------|
| ProfileID        | string  | ProfileID is the unique identifier of a profile. |
| FullName         | string  | Full name                                        |
| UserName         | string  | User name                                        |
| BirthDate        | string  | Birth date                                       |
| Language         | string  | Language code used                               |
| Country          | string  | User country code (VN, US, ID, SG, ...).         |
| Email            | string  | User email address                               |
| EmailVefified    | bool    | True if email is verified, otherwise False       |
| Mobile           | string  | Mobile number                                    |
| Photo            | string  | User profile photo url                           |
| Title            | string  | Title                                            |
| PermanentAddress | Address | Permanent address                                |
| ContactAddress   | Address | Contact address                                  |

##  ProfileAccountInfo

Structure of account information

| Name      | Type   | Description        |
|-----------|--------|--------------------|
| BankCode  | string | Bank code          |
| AccountID | string | Account identifier |
| ProfileID | string | Account identifier |

##  BankCode

DEPRECIATED: BankCode indicates which bank to use; these should be a list of banksthat we are currently integrated withThis is the updated list of banks per January 2018

| Value                        | Description |
|------------------------------|-------------|
| UnknownBankCode              |             |
| Mandiri                      |             |
| Bca                          |             |
| Bni                          |             |
| Bri                          |             |
| DummyBank                    |             |
| Bdo                          |             |
| Bpi                          |             |
| Aceh                         |             |
| Agris                        |             |
| Agroniaga                    |             |
| AmarIndonesia                |             |
| Andara                       |             |
| Antardaerah                  |             |
| AnzIndonesia                 |             |
| Arthagraha                   |             |
| Artos                        |             |
| BangkokIndonesia             |             |
| BarclaysIndonesia            |             |
| BcaSyariah                   |             |
| Bi                           |             |
| BisnisInternasional          |             |
| BniSyariah                   |             |
| BriSyariah                   |             |
| Btpn                         |             |
| BtpnSyariah                  |             |
| Bukopin                      |             |
| BukopinSyariah               |             |
| BumiArta                     |             |
| CapitalIndonesia             |             |
| CentratamaNasional           |             |
| CtbcIndonesia                |             |
| Danamon                      |             |
| Danpac                       |             |
| DinarIndonesia               |             |
| Dki                          |             |
| DkiSyariah                   |             |
| EkonomiRaharja               |             |
| Fama                         |             |
| Ganesha                      |             |
| Harda                        |             |
| IcbcIndonesia                |             |
| InaPerdana                   |             |
| IndexSelindo                 |             |
| Jabar                        |             |
| JabarSyariah                 |             |
| JasaJakarta                  |             |
| Jatim                        |             |
| JatimSyariah                 |             |
| KebHana                      |             |
| KesejahteraanEkonomi         |             |
| MandiriSyariah               |             |
| MaspionIndonesia             |             |
| Mayapada                     |             |
| Maybank                      |             |
| MaybankSyariahIndonesia      |             |
| Mayora                       |             |
| Mega                         |             |
| MegaSyariah                  |             |
| MestikaDharma                |             |
| MetroExpress                 |             |
| Mitraniaga                   |             |
| MizuhoIndonesia              |             |
| MncInternational             |             |
| MuamalatIndonesia            |             |
| MultiArta                    |             |
| Mutiara                      |             |
| NusantaraParahyangan         |             |
| OcbcNisp                     |             |
| OcbcNispSyariah              |             |
| BankOfAmericaIndonesia       |             |
| BankOfChinaIndonesia         |             |
| BankOfIndiaIndonesia         |             |
| Pikko                        |             |
| PrimaMaster                  |             |
| PundiIndonesia               |             |
| QnbIndonesia                 |             |
| RoyalIndonesia               |             |
| SahabatSampoerna             |             |
| SbiIndonesia                 |             |
| SinarHarapan                 |             |
| Sinarmas                     |             |
| Btn                          |             |
| BtnSyariah                   |             |
| UobIndonesia                 |             |
| Victoria                     |             |
| VictoriaSyariah              |             |
| WinduKentjana                |             |
| WooriSaudaraIndonesia        |             |
| YudhaBakti                   |             |
| BnpParibasIndonesia          |             |
| BpdAcehSyariah               |             |
| BpdBali                      |             |
| BpdBengkulu                  |             |
| BpdJambi                     |             |
| BpdJawaTengah                |             |
| BpdKalimantanBarat           |             |
| BpdKalimantanBaratSyariah    |             |
| BpdKalimantanTengah          |             |
| BpdKalimantanTimur           |             |
| BpdKalimantanTimurSyariah    |             |
| BpdKalimantanSelatan         |             |
| BpdKalimantanSelatanSyariah  |             |
| BpdLampung                   |             |
| BpdMaluku                    |             |
| BpdNtb                       |             |
| BpdNtt                       |             |
| BpdPapua                     |             |
| BpdRiauKepri                 |             |
| BpdSulawesiSelatanBarat      |             |
| BpdSulawesiTengah            |             |
| BpdSulawesiTenggara          |             |
| BpdSulawesiUtara             |             |
| BpdSumateraBarat             |             |
| BpdSumateraBaratSyariah      |             |
| BpdSumateraSelatan           |             |
| BpdSumateraSelatanBabel      |             |
| BpdSumateraSelatanSyariah    |             |
| BpdSumateraSelatanUus        |             |
| BpdSumateraUtara             |             |
| BpdSumateraUtaraSyariah      |             |
| BpdYogyakarta                |             |
| BpdYogyakartaSyariah         |             |
| CimbNiaga                    |             |
| CimbNiagaSyariah             |             |
| CitibankIndonesia            |             |
| CommonwealthIndonesia        |             |
| DanamonSyariah               |             |
| DbsIndonesia                 |             |
| DeutscheAgIndonesia          |             |
| NobuIndonesia                |             |
| Panin                        |             |
| PaninSyariah                 |             |
| RaboIndonesia                |             |
| ResonaPerdania               |             |
| StandardCharteredIndonesia   |             |
| SumitomoMitsuiIndonesia      |             |
| BankOfTokyoIndonesia         |             |
| RoyalBankOfScotlandIndonesia |             |
| Permata                      |             |
| PermataSyariah               |             |
| SinarmasSyariah              |             |
| MandiriTaspen                |             |
| Banten                       |             |
| Bei                          |             |
| ChaseIndonesia               |             |
| CcbIndonesia                 |             |
| JtrustIndonesia              |             |
| Transferwise                 |             |

##  CardAccessStatus

| Value                   | Description                                            |
|-------------------------|--------------------------------------------------------|
| UnknownCardAccessStatus |                                                        |
| Often                   | CardAccessStatus_Often indicates a card is used often. |
| Rare                    | CardAccessStatus_Rare indicates a card is used rarely. |

##  CardStatus

| Value             | Description                                        |
|-------------------|----------------------------------------------------|
| UnknownCardStatus |                                                    |
| Lock              | CardStatus_Lock is the lock status for a card.     |
| Unlock            | CardStatus_Unlock is the unlock status for a card. |
| Active            | CardStatus_Active is the active status for a card. |

##  MajorCategory

MajorCategory describes the category of the account.

| Value                | Description                             |
|----------------------|-----------------------------------------|
| UnknownMajorCategory |                                         |
| Dep                  | MajorCategory_Dep for deposit category. |
| Loan                 | MajorCategory_Loan for loan category.   |

##  MajorType

MajorType describes the type of the account.

| Value            | Description                                           |
|------------------|-------------------------------------------------------|
| UnknownMajorType |                                                       |
| Checking         | MajorType_Checking account.                           |
| Saving           | MajorType_Saving account.                             |
| TimeDeposit      | MajorType_TimeDeposit for a time deposit account.     |
| CommercialLoan   | MajorType_CommercialLoan for a business loan account. |
| MortgageLoan     | MajorType_MortgageLoan for a home loan account.       |
| ConsumerLoan     | MajorType_ConsumerLoan for a consumer loan account.   |
