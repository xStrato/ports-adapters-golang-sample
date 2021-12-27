package cli

import (
	"fmt"

	"github.com/xStrato/ports-adapters-go-sample/src/application/interfaces"
)

func Run(s interfaces.IProductService, action, id, name string, price float32) (string, error) {
	result := ""
	switch action {
	case "create":
		p, err := s.Create(name, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product created!! Id: %s, Name: %s, Price: %f, Status: %s", p.GetID(), p.GetName(), p.GetPrice(), p.GetStatus())
	case "enable":
		p, err := s.Get(id)
		if err != nil {
			return result, err
		}
		res, err := s.Enable(p)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been enabled", res.GetName())

	case "disable":
		p, err := s.Get(id)
		if err != nil {
			return result, err
		}
		res, err := s.Disable(p)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been disabled", res.GetName())

	default:
		p, err := s.Get(id)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product Id: %s\n Name: %s\n Price: %f\n Status: %s", p.GetID(), p.GetName(), p.GetPrice(), p.GetStatus())
	}
	return result, nil
}
