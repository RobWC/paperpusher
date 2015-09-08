package main

//Unit basic interface to describe unit behaviors
type Unit interface {
	Attack()
	Defend()
	Flee()
}
