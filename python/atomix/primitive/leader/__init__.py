# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: atomix/primitive/leader/latch.proto
# plugin: python-betterproto
from dataclasses import dataclass
from typing import AsyncIterator, List, Optional

import betterproto
import grpclib


class EventType(betterproto.Enum):
    NONE = 0
    CHANGE = 1


@dataclass(eq=False, repr=False)
class LatchRequest(betterproto.Message):
    headers: "__primitive__.RequestHeaders" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class LatchResponse(betterproto.Message):
    headers: "__primitive__.ResponseHeaders" = betterproto.message_field(1)
    latch: "Latch" = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetRequest(betterproto.Message):
    headers: "__primitive__.RequestHeaders" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetResponse(betterproto.Message):
    headers: "__primitive__.ResponseHeaders" = betterproto.message_field(1)
    latch: "Latch" = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class EventsRequest(betterproto.Message):
    headers: "__primitive__.RequestHeaders" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class EventsResponse(betterproto.Message):
    headers: "__primitive__.ResponseHeaders" = betterproto.message_field(1)
    event: "Event" = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Event(betterproto.Message):
    type: "EventType" = betterproto.enum_field(1)
    latch: "Latch" = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Latch(betterproto.Message):
    meta: "_meta__.ObjectMeta" = betterproto.message_field(1)
    leader: str = betterproto.string_field(2)
    participants: List[str] = betterproto.string_field(3)

    def __post_init__(self) -> None:
        super().__post_init__()


class LeaderLatchServiceStub(betterproto.ServiceStub):
    """Leader latch service"""

    async def latch(
        self, *, headers: "__primitive__.RequestHeaders" = None
    ) -> "LatchResponse":
        """Latch attempts to acquire the leader latch"""

        request = LatchRequest()
        if headers is not None:
            request.headers = headers

        return await self._unary_unary(
            "/atomix.primitive.leader.LeaderLatchService/Latch", request, LatchResponse
        )

    async def get(
        self, *, headers: "__primitive__.RequestHeaders" = None
    ) -> "GetResponse":
        """Get gets the current leader"""

        request = GetRequest()
        if headers is not None:
            request.headers = headers

        return await self._unary_unary(
            "/atomix.primitive.leader.LeaderLatchService/Get", request, GetResponse
        )

    async def events(
        self, *, headers: "__primitive__.RequestHeaders" = None
    ) -> AsyncIterator["EventsResponse"]:
        """Events listens for leader change events"""

        request = EventsRequest()
        if headers is not None:
            request.headers = headers

        async for response in self._unary_stream(
            "/atomix.primitive.leader.LeaderLatchService/Events",
            request,
            EventsResponse,
        ):
            yield response


from .. import meta as _meta__
from ... import primitive as __primitive__
