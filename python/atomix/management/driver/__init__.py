# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: atomix/management/driver/agent.proto, atomix/management/driver/driver.proto
# plugin: python-betterproto
from dataclasses import dataclass
from typing import Optional

import betterproto
import grpclib


@dataclass(eq=False, repr=False)
class ProxyId(betterproto.Message):
    primitive_id: "__primitive__.PrimitiveId" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ProxyOptions(betterproto.Message):
    read: bool = betterproto.bool_field(1)
    write: bool = betterproto.bool_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class CreateProxyRequest(betterproto.Message):
    proxy_id: "ProxyId" = betterproto.message_field(1)
    options: "ProxyOptions" = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class CreateProxyResponse(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class DestroyProxyRequest(betterproto.Message):
    proxy_id: "ProxyId" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class DestroyProxyResponse(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class AgentId(betterproto.Message):
    namespace: str = betterproto.string_field(1)
    name: str = betterproto.string_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class AgentAddress(betterproto.Message):
    host: str = betterproto.string_field(1)
    port: int = betterproto.int32_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class AgentConfig(betterproto.Message):
    protocol: "__protocol__.ProtocolConfig" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class StartAgentRequest(betterproto.Message):
    agent_id: "AgentId" = betterproto.message_field(1)
    address: "AgentAddress" = betterproto.message_field(2)
    config: "AgentConfig" = betterproto.message_field(3)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class StartAgentResponse(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ConfigureAgentRequest(betterproto.Message):
    agent_id: "AgentId" = betterproto.message_field(1)
    config: "AgentConfig" = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ConfigureAgentResponse(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class StopAgentRequest(betterproto.Message):
    agent_id: "AgentId" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class StopAgentResponse(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


class AgentStub(betterproto.ServiceStub):
    async def create_proxy(
        self, *, proxy_id: "ProxyId" = None, options: "ProxyOptions" = None
    ) -> "CreateProxyResponse":

        request = CreateProxyRequest()
        if proxy_id is not None:
            request.proxy_id = proxy_id
        if options is not None:
            request.options = options

        return await self._unary_unary(
            "/atomix.management.driver.Agent/CreateProxy", request, CreateProxyResponse
        )

    async def destroy_proxy(
        self, *, proxy_id: "ProxyId" = None
    ) -> "DestroyProxyResponse":

        request = DestroyProxyRequest()
        if proxy_id is not None:
            request.proxy_id = proxy_id

        return await self._unary_unary(
            "/atomix.management.driver.Agent/DestroyProxy",
            request,
            DestroyProxyResponse,
        )


class DriverStub(betterproto.ServiceStub):
    async def start_agent(
        self,
        *,
        agent_id: "AgentId" = None,
        address: "AgentAddress" = None,
        config: "AgentConfig" = None,
    ) -> "StartAgentResponse":

        request = StartAgentRequest()
        if agent_id is not None:
            request.agent_id = agent_id
        if address is not None:
            request.address = address
        if config is not None:
            request.config = config

        return await self._unary_unary(
            "/atomix.management.driver.Driver/StartAgent", request, StartAgentResponse
        )

    async def configure_agent(
        self, *, agent_id: "AgentId" = None, config: "AgentConfig" = None
    ) -> "ConfigureAgentResponse":

        request = ConfigureAgentRequest()
        if agent_id is not None:
            request.agent_id = agent_id
        if config is not None:
            request.config = config

        return await self._unary_unary(
            "/atomix.management.driver.Driver/ConfigureAgent",
            request,
            ConfigureAgentResponse,
        )

    async def stop_agent(self, *, agent_id: "AgentId" = None) -> "StopAgentResponse":

        request = StopAgentRequest()
        if agent_id is not None:
            request.agent_id = agent_id

        return await self._unary_unary(
            "/atomix.management.driver.Driver/StopAgent", request, StopAgentResponse
        )


from ... import primitive as __primitive__
from ... import protocol as __protocol__