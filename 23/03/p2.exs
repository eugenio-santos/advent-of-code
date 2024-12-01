defmodule D3 do
    def solver do
        lines = File.read!("puzzle") |> String.split("\n") |> List.delete("")

        parseLines(lines)
        |> moses([], [])
        |> gearRatios(lines)
        |> IO.inspect(label: "sol")
    end

    def parseLines(lines) do
        lines
        |> Stream.with_index()
        |> Enum.map(fn l ->
            getNums(l)
        end)
    end

    def getNums(l) do
        {s, i} = l
        x = getN(s, 0, 0, [], [])

        nums = Enum.at(x, 0)
        stars = Enum.at(x, 1)

        nums = Enum.map(nums, & Tuple.insert_at(&1, 0, i))
        stars = Enum.map(stars, & Tuple.insert_at(&1, 0, i))

        {nums, stars}
    end

    def getN("", 0, _, res, stars) do [res, stars] end
    def getN("", i, gc, res, stars) do # end of an number at the end of the line
        [res ++ [{gc-i, i}], stars]
    end
    def getN(l, i,gc, res, stars) when <<48>> <= binary_part(l, 0, 1) and binary_part(l, 0, 1) <= <<58>> do # current char is number
        getN(String.slice(l, 1..-1), i+1, gc+1, res, stars)
    end
    def getN(l, 0, gc, res, stars) when <<42>> == binary_part(l, 0, 1) do # * and not end of a number
        getN(String.slice(l, 1..-1), 0, gc+1, res, stars ++ [{gc, 1}])
    end
    def getN(l, 0, gc, res, stars) do getN(String.slice(l, 1..-1), 0, gc+1, res, stars) end # not a number and not end of a number
    def getN(l, i, gc, res, stars) when <<42>> == binary_part(l, 0, 1) do # * and end of a number
        getN(String.slice(l, 1..-1), 0, gc+1, res ++ [{gc-i, i}], stars ++ [{gc, 1}])
    end
    def getN(l, i, gc, res, stars) do # end of a number
        getN(String.slice(l, 1..-1), 0, gc+1, res ++ [{gc-i, i}], stars)
    end

    def moses([], nums, stars) do {nums, stars} end
    def moses([head | rest], nums, stars) do
        {n, s} = head
        moses(rest, nums ++ [n], stars ++ s)
    end

    def gearRatios({nums, stars}, lines) do
        nums = nums |> Enum.map(fn l ->
            l |>
            Enum.map(fn {y, x, len}  ->
                {y, x, x+len-1}
            end)
        end)

        stars |> Enum.map(fn s ->
            {r, c, _} = s
            getStarNums({r, c}, nums)
        end)
        |> Enum.filter(&(length(&1) == 2))
        |> Enum.reduce(0, fn ns, acc ->
            {y1, x1, l1} = ns |> List.first()
            n1 = Enum.at(lines, y1)
            |> String.slice(x1..l1)
            |> String.to_integer()

            {y2, x2, l2} = ns |> List.last()
            n2 = Enum.at(lines, y2)
            |> String.slice(x2..l2)
            |> String.to_integer()

            acc + n1*n2
        end)
    end

    def getStarNums({0, c}, nums) do
        # validate row (left/right)
        n = validateRow(Enum.at(nums, 0), c)
        # validate bellow
        n ++ validateBellow(Enum.at(nums, 1), c)
    end
    def getStarNums({r, c}, nums) do
        # validate above
        n = validateAbove(Enum.at(nums, r-1), c)
        # validate row (right)
        n = n ++ validateRow(Enum.at(nums, r), c)
        # validate bellow
        n ++ validateBellow(Enum.at(nums, r+1), c)
    end

    def validateAbove(nums, index) do
        nums |> Enum.filter(fn {_, s, e} ->
            (s in (Range.new(index-1,index+1) |> Range.to_list()))
            ||
            (e in (Range.new(index-1,index+1) |> Range.to_list()))
        end)
    end

    def validateRow(nums, index) do
        nums |> Enum.filter(fn {_, s, e} ->
            (s in (Range.new(index-1,index+1) |> Range.to_list()))
            ||
            (e in (Range.new(index-1,index+1) |> Range.to_list()))
        end)
    end

    def validateBellow(nums, index) do
        nums |> Enum.filter(fn {_, s, e} ->
            (s in (Range.new(index-1,index+1) |> Range.to_list()))
            ||
            (e in (Range.new(index-1,index+1) |> Range.to_list()))
        end)
    end


end


D3.solver()
