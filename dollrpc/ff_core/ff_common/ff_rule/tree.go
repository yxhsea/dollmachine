package ff_rule

type RuleNode struct {
	RuleId        int         `json:"m_rule_id"`
	Sort          int         `json:"m_sort"`
	Name          string      `json:"title"`
	NickName      string      `json:"m_nick_name"`
	PRuleId       int         `json:"m_p_rule_id"`
	IsMenuDisplay int         `json:"m_is_menu_display"`
	Path          string      `json:"m_path"`
	Conditions    string      `json:"m_conditions"`
	CreatedAt     int64       `json:"m_created_at"`
	UpdatedAt     int64       `json:"m_updated_at"`
	Status        int         `json:"m_status"`
	Children      []*RuleNode `json:"children"`
}
