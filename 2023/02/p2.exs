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
                Map.merge(m, a)
            end )
        end)
        |>Enum.reduce(%{}, fn m, a ->
            Map.merge(m, a, fn _k, v1, v2 ->
                if v1 >= v2 do
                    v1
                else
                  v2
                end
            end)
        end )
        |>Enum.reduce(1, fn {_, v}, acc ->
            acc * v
        end)
    end

    def s do
        File.stream!("puzzle")
        |> Stream.map(fn line ->
            l = String.trim(line)

            [g | draws] = String.split(l, ":")
            IO.inspect(line)

            gameid = D2.getGameId(g)

            m = List.first(draws)
            |> D2.parseDraw()

        end)
        |> Enum.reduce(fn pow, acc ->
            acc + pow
        end)
    end
end

D2.s()
|>IO.inspect()
