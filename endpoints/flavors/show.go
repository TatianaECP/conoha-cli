package flavors

func Show(tenantId string, tokenId string, name string, item *Flavor) error {
	items := &[]Flavor{}
	err := Get(tenantId, tokenId, items)
	if err != nil {
		return err
	}

	for _, it := range *items {
		if it.Name == name {
			*item = it
			break
		}
	}

	return nil
}
