package activities

type ActivityType int

const (
	TableName = "activities"

	ActivityTypeDefault ActivityType = iota
	ActivityTypeChat
	ActivityTypeCafe
	ActivityTypeBar
	ActivityTypePub
	ActivityTypeWalk
	ActivityTypeHike
)

// The Activity struct
type Activity struct {
	Type     ActivityType
	IconURL  string
	ImageURL string
}

func NewActivity(a ActivityType) *Activity {
	switch a {
	case ActivityTypeDefault:
		return &Activity{
			IconURL:  "/static/icons/camera.svg",
			ImageURL: "/static/images/default.jpg",
		}
	case ActivityTypeChat:
		return &Activity{
			IconURL:  "/static/icons/camera.svg",
			ImageURL: "/static/images/default.jpg",
		}
	case ActivityTypeCafe:
		return &Activity{
			IconURL:  "/static/icons/camera.svg",
			ImageURL: "/static/images/cafe.jpg",
		}
	case ActivityTypeBar:
		return &Activity{
			IconURL:  "/static/icons/camera.svg",
			ImageURL: "/static/images/bar.jpg",
		}
	case ActivityTypePub:
		return &Activity{
			IconURL:  "/static/icons/camera.svg",
			ImageURL: "/static/images/pub.jpg",
		}
	case ActivityTypeWalk:
		return &Activity{
			IconURL:  "/static/icons/camera.svg",
			ImageURL: "/static/images/walk.jpg",
		}
	case ActivityTypeHike:
		return &Activity{
			IconURL:  "/static/icons/camera.svg",
			ImageURL: "/static/images/hike.jpg",
		}
	}
	panic("[NewActivity] Switch failed")
}
