# Launch a server
@async begin
    server = listen(2000)
    while true
        sock = accept(server)
        @async while isopen(sock)
            write(sock, readline(sock))
        end
    end
end

# Connect a TCP client
client = connect(2000)

@async while true
    write(STDOUT, readline(client))
end

println(client, "Hello World!")

# Close the socket
close(client)