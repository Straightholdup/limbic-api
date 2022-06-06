# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import speech_emotion_recognition_pb2 as speech__emotion__recognition__pb2


class SpeechEmotionRecognitionStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.LoadData = channel.stream_unary(
                '/limbic.protos.speech_emotion_recognition.SpeechEmotionRecognition/LoadData',
                request_serializer=speech__emotion__recognition__pb2.Chunk.SerializeToString,
                response_deserializer=speech__emotion__recognition__pb2.LoadDataReply.FromString,
                )


class SpeechEmotionRecognitionServicer(object):
    """Missing associated documentation comment in .proto file."""

    def LoadData(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SpeechEmotionRecognitionServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'LoadData': grpc.stream_unary_rpc_method_handler(
                    servicer.LoadData,
                    request_deserializer=speech__emotion__recognition__pb2.Chunk.FromString,
                    response_serializer=speech__emotion__recognition__pb2.LoadDataReply.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'limbic.protos.speech_emotion_recognition.SpeechEmotionRecognition', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class SpeechEmotionRecognition(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def LoadData(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(request_iterator, target, '/limbic.protos.speech_emotion_recognition.SpeechEmotionRecognition/LoadData',
            speech__emotion__recognition__pb2.Chunk.SerializeToString,
            speech__emotion__recognition__pb2.LoadDataReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
