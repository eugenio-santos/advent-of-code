defmodule D3 do
    def solver do
        lines = File.read!("puzzle") |> String.split("\n") |> List.delete("")
        # lines = File.read!("test") |> String.split("\n") |> List.delete("")
        # lines = File.read!("bug") |> String.split("\n") |> List.delete("")
        |> IO.inspect()

        parseLines(lines)
        |> Enum.filter(& !is_nil(&1))
        |> IO.inspect(label: "numbers with an * adjecent")
        |> Enum.reduce(0, fn n, acc ->
            String.to_integer(n) + acc
        end)
        |> IO.inspect()
    end

    # return list of numbers [[nums], [nums]]
    def parseLines(lines) do
        lines
        |> Stream.with_index()
        |> Enum.map(fn l ->
            getNums(l)
        end)
        |>List.flatten()
        |> IO.inspect()
        |> Enum.map(fn {li, si, len} ->
            # IO.inspect({li, si, len})
            cond do
                validateAbove({li, si, len}, lines) -> String.slice(Enum.at(lines, li), si, len)
                validateRow({si, len}, Enum.at(lines, li)) -> String.slice(Enum.at(lines, li), si, len)
                validateBellow({li, si, len}, Enum.at(lines, li+1)) -> String.slice(Enum.at(lines, li), si, len)
                true -> nil
            end
        end)
    end

    def validateAbove({0, _, _}, _) do false end
    def validateAbove({li, 0, len}, lines) do
        validateAbove({li, 1, len-1}, lines)
    end
    def validateAbove({li, si, len}, lines) do
        # iterate line
        Enum.at(lines, li-1) |> String.slice(si-1, len+2) |> String.match?(~r/[^\d.]/)
    end

    def validateRow({0, len}, line) do
        line |> String.slice(len, 1) |> String.match?(~r/[^\d.]/)
    end
    def validateRow({si, len}, line) do
        line |> String.slice(si-1, 1) |> String.match?(~r/[^\d.]/)
        || line |> String.slice(si+len, 1) |> String.match?(~r/[^\d.]/)
    end


    def validateBellow(_, nil) do false end
    def validateBellow({li, 0, len}, line) do # start of a line
        validateBellow({li, 1, len-1}, line)
    end
    def validateBellow({_, si, len}, line) do
        line |> String.slice(si-1, len+2) |> String.match?(~r/[^\d.]/)
    end

    def sanInd(i) when i < 0 do {:err} end
    def sanInd(i) do {:ok, i} end

    # [{line_i, start_i, len?}]
    def getNums(l) do
        {s, i} = l
        getN(s, 0, 0, [])
        |>Enum.map(fn x ->

            Tuple.insert_at(x, 0, i)
        end)
    end

    # res = [{start_i, len}]
    def getN("", 0, _, res) do res end
    def getN("", i, gc, res) do # end of an number at the end of the line
        res ++ [{gc-i, i}]
    end
    def getN(l, i,gc, res) when <<48>> <= binary_part(l, 0, 1) and binary_part(l, 0, 1) <= <<58>> do # current char is number
        getN(String.slice(l, 1..-1), i+1, gc+1, res)
    end
    def getN(l, 0, gc, res) do getN(String.slice(l, 1..-1), 0, gc+1, res) end # not a number and not end of a number
    def getN(l, i, gc, res) do # end of a number
        getN(String.slice(l, 1..-1), 0, gc+1, res ++ [{gc-i, i}])
    end
end


D3.solver()
