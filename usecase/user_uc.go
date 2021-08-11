package usecase

import "context"

type UserUC struct {
	ctx context.Context
}

func (uc *UserUC) Login(email, password string)  {
	
}
