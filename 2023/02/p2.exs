defmodule D2 do
    def getGameId(s) do
        s
        |> String.split()
        |> List.last()
        |> String.to_integer()
    end

    # %{m | s |> String.split() |> List.last() => s |> String.split() |> List.first()}
    def parseDraw(s) do
        # iterate draws
        s
        |> String.split(";")
        |> Enum.map(fn d ->
            d
            |> String.split(",")
            |> Enum.map(fn c ->
                %{c |> String.split() |> List.last() =>
                c |> String.split() |> List.first() |> String.to_integer()}
            end)
            |>Enum.reduce(%{}, fn m, a ->
                Map.merge(m, a, fn _k, v1, v2 -> v1 + v2 end)
            end )
        end)
        |>Enum.reduce(%{}, fn m, a ->
            Map.merge(m, a, fn _k, v1, v2 -> v1 + v2 end)
        end )
    end

    def s do
        File.stream!("puzzle")
        |> Stream.map(fn line ->
            l = String.trim(line)

            [g | draws] = String.split(l, ":")
            # IO.inspect(draws)

            gameid = D2.getGameId(g)

            m = List.first(draws)
            |> D2.parseDraw()
            # |> IO.inspect()

            cond do
                Map.get(m, "red") > 12 -> {:noGame}
                Map.get(m, "green") > 13 -> {:noGame}
                Map.get(m, "blue") > 14 -> {:noGame}
                true -> {:ok, gameid}
            end
        end)
        |> Enum.reduce(0, fn game, acc ->
            # IO.inspect(game, label: "g")
            IO.inspect(acc, label: "acc")
            case game do
                {:ok, v} -> acc + v
                _ -> acc
            end
        end)
    end
end

D2.s()
|>IO.inspect()
