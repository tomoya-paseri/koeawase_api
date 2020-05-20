// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"local.packages/handler"
	"local.packages/similarity"
	"local.packages/src"
	"local.packages/task"
	"local.packages/voice"
)

// Injectors from wire.go:

func InitializeCLI(client *firestore.Client, ctx context.Context) (*src.CLI, error) {
	repositoryInterface := Voice.NewRepository(client, ctx)
	serviceInterface := Voice.NewService(repositoryInterface)
	similarityRepositoryInterface := Similarity.NewRepository(client, ctx)
	similarityServiceInterface := Similarity.NewService(similarityRepositoryInterface)
	voiceTask := Task.NewVoiceTask(serviceInterface, similarityServiceInterface)
	cli := src.NewCLI(voiceTask)
	return cli, nil
}

func InitializeServer(client *firestore.Client, ctx context.Context) (*src.Server, error) {
	repositoryInterface := Voice.NewRepository(client, ctx)
	serviceInterface := Voice.NewService(repositoryInterface)
	voiceHandler := Handler.NewVoiceHandler(serviceInterface)
	server := src.NewServer(voiceHandler)
	return server, nil
}
