package models

type MasterPassword struct {
	id    string
	Value string `clover:"value"`
	clear string
}

func NewMasterPasswordFromB64(value string) MasterPassword {

	return MasterPassword{
		"",
		value,
		"",
	}
}

func NewMasterPassword(value string) MasterPassword {
	c := newCrypto(value)
	v, _ := c.encryptB64(value)

	return MasterPassword{
		"",
		v,
		"",
	}
}

func (m *MasterPassword) GetCrypto() Crypto {

	return newCrypto(m.clear)
}

func (m *MasterPassword) SetClear(clear string) {
	m.clear = clear
}

/* func (m *MasterPassword) GetClear() string {

	return m.clear
} */

func (m *MasterPassword) Check(value string, cb func(v string)) bool {
	c := newCrypto(value)

	v, err := c.decryptB64(m.Value)
	if err != nil {
		return false
	}

	if v == value {
		cb(value)
		return true
	}

	return false
}
