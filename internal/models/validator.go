package models

func (b Banner) Validate() bool {
	if b.Content == nil || b.FeatureID < 1 {
		return false
	}
	for _, id := range b.TagIDs {
		if id < 1 {
			return false
		}
	}
	return true
}
