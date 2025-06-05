package actioninfo

// ActionInfo интерфейс для работы с информацией о действиях
type ActionInfo interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// ProcessAction обрабатывает действие, используя интерфейс ActionInfo
func ProcessAction(action ActionInfo, data string) (string, error) {
	if err := action.Parse(data); err != nil {
		return "", err
	}
	return action.ActionInfo()
}
