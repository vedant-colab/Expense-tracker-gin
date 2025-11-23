package audit

func (s *auditService) LogSimple(userID, action string) error {
	return s.Log(userID, action, "", "", nil)
}
