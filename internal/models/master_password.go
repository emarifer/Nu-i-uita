package models

type MasterPassword struct {
	Id    string
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

func (m *MasterPassword) Check(value string) bool {
	c := newCrypto(value)

	v, err := c.decryptB64(m.Value)
	if err != nil {
		return false
	}

	if v == value {
		m.clear = value
		return true
	}

	return false
}
