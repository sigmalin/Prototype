using System;
using Grpc.Core;
using Protobufs.helloWorld;
using Protobufs.chat;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class gRpcMain : MonoBehaviour
{
    // Start is called before the first frame update
    void Start()
    {
        helloworld();

        chat();
    }

    void helloworld()
    {
        Channel channel = new Channel("127.0.0.1:5033", ChannelCredentials.Insecure);

        var client = new HelloWorldService.HelloWorldServiceClient(channel);

        var reply = client.Say(new HelloWorldRequest() { Me = "sigma" });
        Debug.Log($"reply = {reply.Msg}");

        channel.ShutdownAsync().Wait();
    }

    async void chat()
    {
        Channel channel = new Channel("127.0.0.1:5033", ChannelCredentials.Insecure);

        var client = new ChatService.ChatServiceClient(channel);

        using(var listener = client.Listen(new ListenRequest() { Name = "sigma" }))
        {
            while (await listener.ResponseStream.MoveNext())
            {
                var result = listener.ResponseStream.Current;
                Debug.Log($"{result.Name} : {result.Msg}");
            }
        }
    }
}
