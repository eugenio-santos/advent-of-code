defmodule D4 do
    def solver() do
        lines = File.read!("puzzle.txt")
        # lines = File.read!("test")
            |> String.split("\n")
            |> List.delete("")
            # |> IO.inspect()

        cardCount = 1..length(lines)
        |> Range.to_list()
        |> Map.new(fn x -> {x, 1} end)
        |> IO.inspect()

        cardScores = lines
        |> parseLines()
        |> Enum.map(& scoreCard(&1))

        IO.inspect(cardScores)

        expandCards(cardCount, 1, cardScores)
        |> IO.inspect(label: "expand")
        |> Map.values()
        |> Enum.reduce(fn x, acc -> acc+x end)
        |> IO.inspect()
    end

    def expandCards(m, _, []) do m end
    def expandCards(m, i, [0 | t]) do expandCards(m, i+1, t) end
    def expandCards(m, i, [h | t]) do
        multi = Map.fetch!(m, i)
        m2 = Range.to_list(i+1..i+h)
        |> Map.new(fn x ->
            {x, Map.fetch!(m, x)+(1*multi)}
        end)

        expandCards(Map.merge(m, m2), i+1, t)

    end


    def parseLines(lines) do
      lines
      |> Enum.map(fn l ->
        l
        |> String.split(":")
        |> List.last()
        |> String.split("|")
        |> Enum.map(&(String.split(&1)))
      end)
    end

    def scoreCard(card) do
        w = List.first(card)
        nums = List.last(card)

        nums
        |> Enum.reduce(0, fn n, acc ->
            if n in w do
                acc+1
            else
                acc
            end
        end)
    end
end

D4.solver()
