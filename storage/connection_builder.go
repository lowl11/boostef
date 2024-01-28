package storage

import "time"

type Builder struct {
	connectionString string

	maxConnections     int
	maxIdleConnections int

	maxLifetime     time.Duration
	maxIdleLifetime time.Duration
}

func NewBuilder(connectionString string) *Builder {
	return &Builder{
		connectionString: connectionString,
	}
}

func (builder *Builder) SetMaxConnections(maxConnections int) *Builder {
	builder.maxConnections = maxConnections
	return builder
}

func (builder *Builder) SetIdleMaxConnections(maxIdleConnections int) *Builder {
	builder.maxIdleConnections = maxIdleConnections
	return builder
}

func (builder *Builder) SetMaxLifetime(lifetime time.Duration) *Builder {
	builder.maxLifetime = lifetime
	return builder
}

func (builder *Builder) SetIdleMaxLifetime(lifetime time.Duration) *Builder {
	builder.maxIdleLifetime = lifetime
	return builder
}
