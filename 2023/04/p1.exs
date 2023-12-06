defmodule D4 do
    def solver() do
        lines = File.read!("puzzle.txt")
        # lines = File.read!("test")
            |> String.split("\n")
            |> List.delete("")

        lines
        |> parseLines()
        |> Enum.map(& scoreCard(&1))
        |> Enum.filter(& &1!=0)
        |> Enum.reduce(0, fn x, acc ->
           acc + (:math.pow(2, x-1))
        end)
        |> IO.inspect()
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
