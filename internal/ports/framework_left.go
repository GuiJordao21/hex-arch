package ports

import (
	// "context"
)

type GRPCPort interface{
	Run()
	GetAddition()
	GetSubtraction()
	GetMultiplication()
	GetDivision()
}