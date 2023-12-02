IO.puts("hyea")

File.stream!("test")
|> Stream.map(fn line ->
  l = String.trim(line)
  IO.puts(l)

  [game | draws] = String.split(l, ":")
  IO.puts(draws)

  # get game id
  gameid =
    game
    |> String.split()
    |> List.last()
    |> String.to_integer()

  List.first(draws)
  |> String.trim()
  |> String.split(";")
  # iterate dr
  |> Enum.map(fn d ->
    IO.puts(d)
  end)
end)
|> Stream.run()
