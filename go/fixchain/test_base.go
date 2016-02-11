package fixchain

import (
	"fmt"
	"encoding/pem"
	"errors"
	"strings"
	"testing"

	"github.com/google/certificate-transparency/go/x509"
	"github.com/google/certificate-transparency/go/x509/pkix"
)

// CertificateFromPEM takes a string representing a certificate in PEM format
// and returns the corresponding x509.Certificate object
func CertificateFromPEM(pemBytes string) (*x509.Certificate, error) {
	block, _ := pem.Decode([]byte(pemBytes))
	if block == nil {
		return nil, errors.New("failed to decode PEM")
	}
	return x509.ParseCertificate(block.Bytes)
}

// GetTestCertificateFromPEM returns an x509.Certificate from a certificate in
// PEM format for testing purposes.  Any errors in the PEM decoding process are
// reported to the testing framework
func GetTestCertificateFromPEM(t *testing.T, pemBytes string) *x509.Certificate {
	cert, err := CertificateFromPEM(pemBytes)
	if err != nil {
		t.Errorf("Failed to parse leaf: %s", err)
	}
	return cert
}

func nameToKey(name *pkix.Name) string {
	return fmt.Sprintf("%s/%s/%s/%s", strings.Join(name.Country, ","),
		strings.Join(name.Organization, ","),
		strings.Join(name.OrganizationalUnit, ","), name.CommonName)
}

const verisignRoot = `-----BEGIN CERTIFICATE-----
MIICPDCCAaUCEHC65B0Q2Sk0tjjKewPMur8wDQYJKoZIhvcNAQECBQAwXzELMAkG
A1UEBhMCVVMxFzAVBgNVBAoTDlZlcmlTaWduLCBJbmMuMTcwNQYDVQQLEy5DbGFz
cyAzIFB1YmxpYyBQcmltYXJ5IENlcnRpZmljYXRpb24gQXV0aG9yaXR5MB4XDTk2
MDEyOTAwMDAwMFoXDTI4MDgwMTIzNTk1OVowXzELMAkGA1UEBhMCVVMxFzAVBgNV
BAoTDlZlcmlTaWduLCBJbmMuMTcwNQYDVQQLEy5DbGFzcyAzIFB1YmxpYyBQcmlt
YXJ5IENlcnRpZmljYXRpb24gQXV0aG9yaXR5MIGfMA0GCSqGSIb3DQEBAQUAA4GN
ADCBiQKBgQDJXFme8huKARS0EN8EQNvjV69qRUCPhAwL0TPZ2RHP7gJYHyX3KqhE
BarsAx94f56TuZoAqiN91qyFomNFx3InzPRMxnVx0jnvT0Lwdd8KkMaOIG+YD/is
I19wKTakyYbnsZogy1Olhec9vn2a/iRFM9x2Fe0PonFkTGUugWhFpwIDAQABMA0G
CSqGSIb3DQEBAgUAA4GBALtMEivPLCYATxQT3ab7/AoRhIzzKBxnki98tsX63/Do
lbwdj2wsqFHMc9ikwFPwTtYmwHYBV4GSXiHx0bH/59AhWM1pF+NEHJwZRDmJXNyc
AA9WjQKZ7aKQRUzkuxCkPfAyAw7xzvjoyVGM5mKf5p/AfbdynMk2OmufTqj/ZA1k
-----END CERTIFICATE-----
`

const thawteIntermediate = `-----BEGIN CERTIFICATE-----
MIIDIzCCAoygAwIBAgIEMAAAAjANBgkqhkiG9w0BAQUFADBfMQswCQYDVQQGEwJV
UzEXMBUGA1UEChMOVmVyaVNpZ24sIEluYy4xNzA1BgNVBAsTLkNsYXNzIDMgUHVi
bGljIFByaW1hcnkgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMDQwNTEzMDAw
MDAwWhcNMTQwNTEyMjM1OTU5WjBMMQswCQYDVQQGEwJaQTElMCMGA1UEChMcVGhh
d3RlIENvbnN1bHRpbmcgKFB0eSkgTHRkLjEWMBQGA1UEAxMNVGhhd3RlIFNHQyBD
QTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA1NNn0I0Vf67NMf59HZGhPwtx
PKzMyGT7Y/wySweUvW+Aui/hBJPAM/wJMyPpC3QrccQDxtLN4i/1CWPN/0ilAL/g
5/OIty0y3pg25gqtAHvEZEo7hHUD8nCSfQ5i9SGraTaEMXWQ+L/HbIgbBpV8yeWo
3nWhLHpo39XKHIdYYBkCAwEAAaOB/jCB+zASBgNVHRMBAf8ECDAGAQH/AgEAMAsG
A1UdDwQEAwIBBjARBglghkgBhvhCAQEEBAMCAQYwKAYDVR0RBCEwH6QdMBsxGTAX
BgNVBAMTEFByaXZhdGVMYWJlbDMtMTUwMQYDVR0fBCowKDAmoCSgIoYgaHR0cDov
L2NybC52ZXJpc2lnbi5jb20vcGNhMy5jcmwwMgYIKwYBBQUHAQEEJjAkMCIGCCsG
AQUFBzABhhZodHRwOi8vb2NzcC50aGF3dGUuY29tMDQGA1UdJQQtMCsGCCsGAQUF
BwMBBggrBgEFBQcDAgYJYIZIAYb4QgQBBgpghkgBhvhFAQgBMA0GCSqGSIb3DQEB
BQUAA4GBAFWsY+reod3SkF+fC852vhNRj5PZBSvIG3dLrWlQoe7e3P3bB+noOZTc
q3J5Lwa/q4FwxKjt6lM07e8eU9kGx1Yr0Vz00YqOtCuxN5BICEIlxT6Ky3/rbwTR
bcV0oveifHtgPHfNDs5IAn8BL7abN+AqKjbc1YXWrOU/VG+WHgWv
-----END CERTIFICATE-----
`

const googleLeaf = `-----BEGIN CERTIFICATE-----
MIIDITCCAoqgAwIBAgIQL9+89q6RUm0PmqPfQDQ+mjANBgkqhkiG9w0BAQUFADBM
MQswCQYDVQQGEwJaQTElMCMGA1UEChMcVGhhd3RlIENvbnN1bHRpbmcgKFB0eSkg
THRkLjEWMBQGA1UEAxMNVGhhd3RlIFNHQyBDQTAeFw0wOTEyMTgwMDAwMDBaFw0x
MTEyMTgyMzU5NTlaMGgxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlh
MRYwFAYDVQQHFA1Nb3VudGFpbiBWaWV3MRMwEQYDVQQKFApHb29nbGUgSW5jMRcw
FQYDVQQDFA53d3cuZ29vZ2xlLmNvbTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkC
gYEA6PmGD5D6htffvXImttdEAoN4c9kCKO+IRTn7EOh8rqk41XXGOOsKFQebg+jN
gtXj9xVoRaELGYW84u+E593y17iYwqG7tcFR39SDAqc9BkJb4SLD3muFXxzW2k6L
05vuuWciKh0R73mkszeK9P4Y/bz5RiNQl/Os/CRGK1w7t0UCAwEAAaOB5zCB5DAM
BgNVHRMBAf8EAjAAMDYGA1UdHwQvMC0wK6ApoCeGJWh0dHA6Ly9jcmwudGhhd3Rl
LmNvbS9UaGF3dGVTR0NDQS5jcmwwKAYDVR0lBCEwHwYIKwYBBQUHAwEGCCsGAQUF
BwMCBglghkgBhvhCBAEwcgYIKwYBBQUHAQEEZjBkMCIGCCsGAQUFBzABhhZodHRw
Oi8vb2NzcC50aGF3dGUuY29tMD4GCCsGAQUFBzAChjJodHRwOi8vd3d3LnRoYXd0
ZS5jb20vcmVwb3NpdG9yeS9UaGF3dGVfU0dDX0NBLmNydDANBgkqhkiG9w0BAQUF
AAOBgQCfQ89bxFApsb/isJr/aiEdLRLDLE5a+RLizrmCUi3nHX4adpaQedEkUjh5
u2ONgJd8IyAPkU0Wueru9G2Jysa9zCRo1kNbzipYvzwY4OA8Ys+WAi0oR1A04Se6
z5nRUP8pJcA2NhUzUnC+MY+f6H/nEQyNv4SgQhqAibAxWEEHXw==
-----END CERTIFICATE-----
`

const smimeLeaf = `-----BEGIN CERTIFICATE-----
MIIFBjCCA+6gAwIBAgISESFvrjT8XcJTEe6rBlPptILlMA0GCSqGSIb3DQEBBQUA
MFQxCzAJBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9iYWxTaWduIG52LXNhMSowKAYD
VQQDEyFHbG9iYWxTaWduIFBlcnNvbmFsU2lnbiAyIENBIC0gRzIwHhcNMTIwMTIz
MTYzNjU5WhcNMTUwMTIzMTYzNjU5WjCBlDELMAkGA1UEBhMCVVMxFjAUBgNVBAgT
DU5ldyBIYW1zcGhpcmUxEzARBgNVBAcTClBvcnRzbW91dGgxGTAXBgNVBAoTEEds
b2JhbFNpZ24sIEluYy4xEzARBgNVBAMTClJ5YW4gSHVyc3QxKDAmBgkqhkiG9w0B
CQEWGXJ5YW4uaHVyc3RAZ2xvYmFsc2lnbi5jb20wggEiMA0GCSqGSIb3DQEBAQUA
A4IBDwAwggEKAoIBAQC4ASSTvavmsFQAob60ukSSwOAL9nT/s99ltNUCAf5fPH5j
NceMKxaQse2miOmRRIXaykcq1p/TbI70Ztce38r2mbOwqDHHPVi13GxJEyUXWgaR
BteDMu5OGyWNG1kchVsGWpbstT0Z4v0md5m1BYFnxB20ebJyOR2lXDxsFK28nnKV
+5eMj76U8BpPQ4SCH7yTMG6y0XXsB3cCrBKr2o3TOYgEKv+oNnbaoMt3UxMt9nSf
9jyIshjqfnT5Aew3CUNMatO55g5FXXdIukAweg1YSb1ls05qW3sW00T3d7dQs9/7
NuxCg/A2elmVJSoy8+MLR8JSFEf/aMgjO/TyLg/jAgMBAAGjggGPMIIBizAOBgNV
HQ8BAf8EBAMCBaAwTQYDVR0gBEYwRDBCBgorBgEEAaAyASgKMDQwMgYIKwYBBQUH
AgEWJmh0dHBzOi8vd3d3Lmdsb2JhbHNpZ24uY29tL3JlcG9zaXRvcnkvMCQGA1Ud
EQQdMBuBGXJ5YW4uaHVyc3RAZ2xvYmFsc2lnbi5jb20wCQYDVR0TBAIwADAdBgNV
HSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwQwQwYDVR0fBDwwOjA4oDagNIYyaHR0
cDovL2NybC5nbG9iYWxzaWduLmNvbS9ncy9nc3BlcnNvbmFsc2lnbjJnMi5jcmww
VQYIKwYBBQUHAQEESTBHMEUGCCsGAQUFBzAChjlodHRwOi8vc2VjdXJlLmdsb2Jh
bHNpZ24uY29tL2NhY2VydC9nc3BlcnNvbmFsc2lnbjJnMi5jcnQwHQYDVR0OBBYE
FFWiECe0/L72eVYqcWYnLV6SSjzhMB8GA1UdIwQYMBaAFD8V0m18L+cxnkMKBqiU
bCw7xe5lMA0GCSqGSIb3DQEBBQUAA4IBAQAhQi6hLPeudmf3IBF4IDzCvRI0FaYd
BKfprSk/H0PDea4vpsLbWpA0t0SaijiJYtxKjlM4bPd+2chb7ejatDdyrZIzmDVy
q4c30/xMninGKokpYA11/Ve+i2dvjulu65qasrtQRGybAuuZ67lrp/K3OMFgjV5N
C3AHYLzvNU4Dwc4QQ1BaMOg6KzYSrKbABRZajfrpC9uiePsv7mDIXLx/toBPxWNl
a5vJm5DrZdn7uHdvBCE6kMykbOLN5pmEK0UIlwKh6Qi5XD0pzlVkEZliFkBMJgub
d/eF7xeg7TKPWC5xyOFp9SdMolJM7LTC3wnSO3frBAev+q/nGs9Xxyvs
-----END CERTIFICATE-----
`
