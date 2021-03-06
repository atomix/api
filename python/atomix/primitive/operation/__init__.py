# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: atomix/primitive/operation.proto
# plugin: python-betterproto
from dataclasses import dataclass

import betterproto


class OperationType(betterproto.Enum):
    """OperationType is an enum for specifying the type of operation"""

    COMMAND = 0
    QUERY = 1


class AggregateStrategy(betterproto.Enum):
    """
    AggregateStrategy is an enum for indicating the strategy used to aggregate
    a field
    """

    CHOOSE_FIRST = 0
    APPEND = 1
    SUM = 2
