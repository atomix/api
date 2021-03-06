# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: atomix/primitive/set/set.proto
# plugin: python-betterproto
from dataclasses import dataclass
from typing import AsyncIterator, Optional

import betterproto
import grpclib


class EventType(betterproto.Enum):
    NONE = 0
    ADD = 1
    REMOVE = 2
    REPLAY = 3


@dataclass(eq=False, repr=False)
class SizeRequest(betterproto.Message):
    headers: "__primitive__.RequestHeaders" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class SizeResponse(betterproto.Message):
    headers: "__primitive__.ResponseHeaders" = betterproto.message_field(1)
    size: int = betterproto.uint32_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ContainsRequest(betterproto.Message):
    headers: "__primitive__.RequestHeaders" = betterproto.message_field(1)
    element: "Element" = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ContainsResponse(betterproto.Message):
    headers: "__primitive__.ResponseHeaders" = betterproto.message_field(1)
    contains: bool = betterproto.bool_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class AddRequest(betterproto.Message):
    headers: "__primitive__.RequestHeaders" = betterproto.message_field(1)
    element: "Element" = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class AddResponse(betterproto.Message):
    headers: "__primitive__.ResponseHeaders" = betterproto.message_field(1)
    element: "Element" = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class RemoveRequest(betterproto.Message):
    headers: "__primitive__.RequestHeaders" = betterproto.message_field(1)
    element: "Element" = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class RemoveResponse(betterproto.Message):
    headers: "__primitive__.ResponseHeaders" = betterproto.message_field(1)
    element: "Element" = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ClearRequest(betterproto.Message):
    headers: "__primitive__.RequestHeaders" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ClearResponse(betterproto.Message):
    headers: "__primitive__.ResponseHeaders" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class EventsRequest(betterproto.Message):
    headers: "__primitive__.RequestHeaders" = betterproto.message_field(1)
    replay: bool = betterproto.bool_field(2)

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
    element: "Element" = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ElementsRequest(betterproto.Message):
    headers: "__primitive__.RequestHeaders" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ElementsResponse(betterproto.Message):
    headers: "__primitive__.ResponseHeaders" = betterproto.message_field(1)
    element: "Element" = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Element(betterproto.Message):
    meta: "_meta__.ObjectMeta" = betterproto.message_field(1)
    value: str = betterproto.string_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


class SetServiceStub(betterproto.ServiceStub):
    """Set service"""

    async def size(
        self, *, headers: "__primitive__.RequestHeaders" = None
    ) -> "SizeResponse":
        """Size gets the number of elements in the set"""

        request = SizeRequest()
        if headers is not None:
            request.headers = headers

        return await self._unary_unary(
            "/atomix.primitive.set.SetService/Size", request, SizeResponse
        )

    async def contains(
        self,
        *,
        headers: "__primitive__.RequestHeaders" = None,
        element: "Element" = None,
    ) -> "ContainsResponse":
        """Contains returns whether the set contains a value"""

        request = ContainsRequest()
        if headers is not None:
            request.headers = headers
        if element is not None:
            request.element = element

        return await self._unary_unary(
            "/atomix.primitive.set.SetService/Contains", request, ContainsResponse
        )

    async def add(
        self,
        *,
        headers: "__primitive__.RequestHeaders" = None,
        element: "Element" = None,
    ) -> "AddResponse":
        """Add adds a value to the set"""

        request = AddRequest()
        if headers is not None:
            request.headers = headers
        if element is not None:
            request.element = element

        return await self._unary_unary(
            "/atomix.primitive.set.SetService/Add", request, AddResponse
        )

    async def remove(
        self,
        *,
        headers: "__primitive__.RequestHeaders" = None,
        element: "Element" = None,
    ) -> "RemoveResponse":
        """Remove removes a value from the set"""

        request = RemoveRequest()
        if headers is not None:
            request.headers = headers
        if element is not None:
            request.element = element

        return await self._unary_unary(
            "/atomix.primitive.set.SetService/Remove", request, RemoveResponse
        )

    async def clear(
        self, *, headers: "__primitive__.RequestHeaders" = None
    ) -> "ClearResponse":
        """Clear removes all values from the set"""

        request = ClearRequest()
        if headers is not None:
            request.headers = headers

        return await self._unary_unary(
            "/atomix.primitive.set.SetService/Clear", request, ClearResponse
        )

    async def events(
        self, *, headers: "__primitive__.RequestHeaders" = None, replay: bool = False
    ) -> AsyncIterator["EventsResponse"]:
        """Events listens for set change events"""

        request = EventsRequest()
        if headers is not None:
            request.headers = headers
        request.replay = replay

        async for response in self._unary_stream(
            "/atomix.primitive.set.SetService/Events",
            request,
            EventsResponse,
        ):
            yield response

    async def elements(
        self, *, headers: "__primitive__.RequestHeaders" = None
    ) -> AsyncIterator["ElementsResponse"]:
        """Elements lists all elements in the set"""

        request = ElementsRequest()
        if headers is not None:
            request.headers = headers

        async for response in self._unary_stream(
            "/atomix.primitive.set.SetService/Elements",
            request,
            ElementsResponse,
        ):
            yield response


from .. import meta as _meta__
from ... import primitive as __primitive__
