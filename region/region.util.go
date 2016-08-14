package region

import "errors"

func (r *Region) Validate() error {
	if r.Id == "" {
		return errors.New("ID must not be null.")
	}
	if r.Name == "" {
		return errors.New("Name must not be null.")
	}
	if r.IpRange == nil {
		return errors.New("IP range must not be null.")
	}
	if err := r.IpRange.Validate(); err != nil {
		return err
	}
	if r.State == nil {
		return errors.New("State must not be null.")
	}
	if r.Location == nil {
		return errors.New("Location must not be null.")
	}
	if err := r.Location.Validate(); err != nil {
		return err
	}
	return nil
}
