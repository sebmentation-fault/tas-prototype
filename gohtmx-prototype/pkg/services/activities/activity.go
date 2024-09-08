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
//
// Contains an enum for which type it is, as well as urls to it's icon and
// image.
type Activity struct {
	Type     ActivityType `json:"type"`
	Name     string       `json:"name"`
	IconURL  string       `json:"icon_url"`
	ImageURL string       `json:"image_url"`
}

func NewActivity(a ActivityType) *Activity {
	switch a {
	case ActivityTypeChat:
		return &Activity{
			Type:     ActivityTypeDefault,
			Name:     "Take a selfie",
			IconURL:  "/static/icons/camera.svg",
			ImageURL: "/static/images/default.jpg",
		}
	case ActivityTypeCafe:
		return &Activity{
			Type:     ActivityTypeDefault,
			Name:     "Take a selfie",
			IconURL:  "/static/icons/camera.svg",
			ImageURL: "/static/images/cafe.jpg",
		}
	case ActivityTypeBar:
		return &Activity{
			Type:     ActivityTypeBar,
			Name:     "Drink at bar",
			IconURL:  "/static/icons/camera.svg",
			ImageURL: "/static/images/bar.jpg",
		}
	case ActivityTypePub:
		return &Activity{
			Type:     ActivityTypePub,
			Name:     "Drink at pub",
			IconURL:  "/static/icons/camera.svg",
			ImageURL: "/static/images/pub.jpg",
		}
	case ActivityTypeWalk:
		return &Activity{
			Type:     ActivityTypeWalk,
			Name:     "Go for a walk",
			IconURL:  "/static/icons/camera.svg",
			ImageURL: "/static/images/walk.jpg",
		}
	case ActivityTypeHike:
		return &Activity{
			Type:     ActivityTypeDefault,
			Name:     "Go for a hike",
			IconURL:  "/static/icons/camera.svg",
			ImageURL: "/static/images/hike.jpg",
		}
	}
	return &Activity{
		Type:     ActivityTypeDefault,
		Name:     "Take a selfie",
		IconURL:  "/static/icons/camera.svg",
		ImageURL: "/static/images/default.jpg",
	}
}
