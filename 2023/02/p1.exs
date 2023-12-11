defmodule D2 do
  def getGameId(s) do
    s
    |> String.split()
    |> List.last()
    |> String.to_integer()
  end

  def parseGame(s) do
    s
    |> String.split(";")
    |> Enum.map(fn d -> D2.parseDraw(d) end)
  end

  def parseDraw(d) do
    d
    |> String.split(",")
    |> Enum.map(fn c ->
      # %{c |> String.split() |> List.last() =>
      # c |> String.split() |> List.first() |> String.to_integer()}
      a = String.split(c)
      [h, t] = a
      [String.to_integer(h), t]
    end)
    |> IO.inspect(label: "map draw")
    |> Enum.all?(fn m ->
      m
      |> IO.inspect(label: "m")
      |> D2.colorPicker()
      |> IO.inspect()
    end)

    # |>IO.inspect()
  end

  def colorPicker([v, "red"]) when v <= 12, do: true
  def colorPicker([v, "green"]) when v <= 13, do: true
  def colorPicker([v, "blue"]) when v <= 14, do: true
  def colorPicker([_, _]), do: false

  def s do
    File.stream!("puzzle")
    |> Stream.map(fn line ->
      l = String.trim(line)

      [g | draws] = String.split(l, ":")
      IO.inspect(line, label: "game")

      gameid = D2.getGameId(g)

      m =
        List.first(draws)
        |> D2.parseGame()
        |> Enum.all?()

      # |> IO.inspect()

      if m do
        {:ok, gameid}
      else
        {:nogame}
      end
    end)
    |> Enum.reduce(0, fn game, acc ->
      IO.inspect(game, label: "g")
      # IO.inspect(acc, label: "acc")
      case game do
        {:ok, v} -> acc + v
        _ -> acc
      end
    end)
  end
end

D2.s()
|> IO.inspect()
