SMOKE_PROTO_DIR=local_net/test/smoke/thornode_proto
SMOKE_DOCKER_OPTS = --network=host --rm -e RUNE=THOR.RUNE -e LOGLEVEL=INFO -e PYTHONPATH=/app -w /app -v ${PWD}/local_net/test/smoke:/app

cli-mocknet:
	@docker compose -f local_net/build/docker/docker-compose.yml run --rm cli

run-mocknet:
	@docker compose -f local_net/build/docker/docker-compose.yml --profile mocknet --profile midgard up -d

stop-mocknet:
	@docker compose -f local_net/build/docker/docker-compose.yml --profile mocknet --profile midgard down -v

bootstrap-mocknet: $(SMOKE_PROTO_DIR)
	@docker run ${SMOKE_DOCKER_OPTS} \
		registry.gitlab.com/thorchain/thornode:smoke \
		python scripts/smoke.py --bootstrap-only=True