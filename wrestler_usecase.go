package main

import ()

type wrestlerUsecase struct {
	wrestlerRepository WrestlerRepository
}

func (w wrestlerUsecase) Sign(wrestler *Wrestler) error{
	
	return nil
}