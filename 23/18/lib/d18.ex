defmodule D18 do
    def s() do
        lines = File.read!("puzzle")
        # lines = File.read!("test")
        # lines = File.read!("debug")
            |> String.split("\n")
            |> List.delete("")

        ix = 1
        iy = 1
        {cords, minmax} = buildCords(lines, ix, iy, {nil, 1, nil, 1}, [{1, 1}])
        IO.inspect(minmax, label: "{minr, maxr, minc, maxc")

        normalized_cord = cords
        |> Enum.map(fn c -> normalize(c, minmax) end)
        # |> IO.inspect()

        {rowNum, colNum} = getMatrixD(minmax)

        matrix = Matrex.zeros({rowNum, colNum})

        out = outline(matrix, normalized_cord, nil)

        IO.inspect(Matrex.to_list_of_lists(Matrex.submatrix(out, 50..100, 50..100)), limit: :infinity, width: :infinity )
        Matrex.sum(out)


        out = fillRow(out, 1, rowNum, 1, colNum, false, false)
        # Matrex.sum()
        # IO.inspect(out)
        IO.inspect(Matrex.to_list_of_lists(Matrex.submatrix(out, 50..100, 50..100)), limit: :infinity, width: :infinity )

        Matrex.sum(out)
    end

    def fillRow(matrix, row, maxRow, col, maxCol, _, _) when row >= maxRow, do: matrix
    def fillRow(matrix, row, maxRow, col, maxCol, _, _) when col >= maxCol, do: fillRow(matrix, row+1, maxRow, 1, maxCol, false, false)
    def fillRow(matrix, row, maxRow, col, maxCol, isInside, wasOne) do
        # IO.inspect({row, col, isInside, wasOne})
        if matrix[row][col] == 1.0 and wasOne do
            fillRow(matrix, row, maxRow, col+1, maxCol, isInside, wasOne)
        else
            if matrix[row][col] == 1.0 and !wasOne do
                fillRow(matrix, row, maxRow, col+1, maxCol, !isInside, !wasOne)
            else
                if isInside do
                    fillRow(Matrex.set(matrix, row, col, 1), row, maxRow, col+1, maxCol, isInside, false)
                else
                    fillRow(matrix, row, maxRow, col+1, maxCol, isInside, false)
                end
            end
        end
    end



    def outline(matrix, [], _), do: matrix
    def outline(matrix, [{r, c} | cords], nil), do: outline(Matrex.set(matrix, r, c, 1), cords, {r, c} )

    def outline(matrix, [curr | cords], prev) when curr == prev do
        {r, c} = curr
        outline(Matrex.set(matrix, r, c, 1), cords, prev)
    end

    def outline(matrix, [{r, c} | cords], {pr, pc}) when r == pr and c != pc do
        pc = pc + round((c - pc)/abs(c-pc))
        # IO.inspect({{r, c} , {pr, pc}}, label: "column")
        outline(Matrex.set(matrix, pr, pc, 1), [{r, c} | cords], {pr, pc})
    end
    def outline(matrix, [{r, c} | cords], {pr, pc}) when c == pc and r != pr do
        pr = pr + round((r - pr)/abs(r-pr))
        # IO.inspect({{r, c} , {pr, pc}}, label: "row")
        outline(Matrex.set(matrix, pr, pc, 1), [{r, c} | cords], {pr, pc})
    end

    def normalize({r, c} , {minr, _, minc, _}) when minr <= 0 and minc <= 0 do {r + abs(minr)+1, c + abs(minc)+1} end
    def normalize({r, c} , {minr, _, _, _}) when minr <= 0 do {r + abs(minr)+1, c } end
    def normalize({r, c} , {_, _, minc, _}) when minc <= 0 do {r, c + abs(minc)+1} end
    def normalize({r, c} , {_, _, _, _}) do {r, c} end


    def buildCords([], _, _, minmax ,c), do: {c, minmax}
    def buildCords([mov | tail], r, c, {minr, maxr, minc, maxc}, cords) do

        {nr, nc} = String.split(mov)
        |> updatedCord(r, c)

        # IO.inspect({mov, {nr, nc}})

        buildCords(tail, nr, nc, {
                min(minr, nr),
                max(maxr, nr),
                min(minc, nc),
                max(maxc, nc)
            },
            cords ++ [{nr, nc}])
    end

    def updatedCord(["R", m, _], r, c), do: {r, c+String.to_integer(m)}
    def updatedCord(["D", m, _], r, c), do: {r+String.to_integer(m), c}
    def updatedCord(["U", m, _], r, c), do: {r-String.to_integer(m), c}
    def updatedCord(["L", m, _], r, c), do: {r, c-String.to_integer(m)}


    def getMatrixD({minr, maxr, minc, maxc}) do
        minmaxa = [{minr, minc}, {maxr, maxc}]

        [{minr, minc}, {maxr, maxc}] = minmaxa
        |> Enum.map(fn c -> normalize(c, {minr, maxr, minc, maxc}) end)
        |> IO.inspect(label: "matrix Dim")

        {maxr-minr+1, maxc-minc+1}
    end
end
