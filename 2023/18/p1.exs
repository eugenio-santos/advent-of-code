defmodule D18 do
    def solver() do
        # lines = File.read!("puzzle")
        # lines = File.read!("test")
        lines = File.read!("debug")
            |> String.split("\n")
            |> List.delete("")

        ix = 1
        iy = 0
        cords = buildCords(lines, ix, iy, [{0, 0}])
        |> IO.inspect()

        # [f | _] = cords
        # IO.inspect(cords++[f])

        shoelace(cords, nil, nil, 0)
        # shoelace([ {0, 0}, {7, 0}, {7, -5}, {5, -5}, {5, -7}, {7, -7}, {7, -9}, {2, -9}, {2, -7}, {1, -7}, {1, -5}, {3, -5}, {3, -2}, {1, -2}, {0, 0} ], nil, nil, 0)
        # shoelace([{0, 0}, {4, 0}, {4, -4}, {0, -4}, {0, 0}], nil, nil, 0)
                # [{0, 0}, {4, 1}, {4, -3}, {0, -3}, {0, 1}]
        |> IO.inspect()

    end

    def shoelace([], _, _, t), do: abs(t/2)
    def shoelace([f | tail], nil, nil , t), do: shoelace(tail, f, nil, t)
    def shoelace([s | tail], f, nil , t), do: shoelace(tail, f, s, t)
    def shoelace([h | tail], {x1, y1}, {x2, y2}, t) do

        it = (x1*y2) - (x2*y1)

        shoelace(tail, {x2, y2}, h, t+it)
    end

    def buildCords([], _, _, c), do: c
    # def buildCords([], _, _, _), do: "something went wrong"
    def buildCords([mov | tail], x, y, cords) do

        {nx, ny} = String.split(mov)
        |> updatedCord(x, y)

        IO.inspect({mov, {nx, ny}})

        buildCords(tail, nx, ny, cords ++ [{nx, ny}])
    end

    def updatedCord(["R", m, _], x, y), do: {x+String.to_integer(m), y}
    def updatedCord(["D", m, _], x, y), do: {x, y-String.to_integer(m)}
    def updatedCord(["U", m, _], x, y), do: {x, y+String.to_integer(m)}
    def updatedCord(["L", m, _], x, y), do: {x-String.to_integer(m), y}

end

D18.solver()
