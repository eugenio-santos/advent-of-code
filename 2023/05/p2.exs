defmodule D5 do
    def solver() do
        lines = File.read!("puzzle")
        # lines = File.read!("test")
            |> String.split("\n")

        [seeds | lines] = lines

        seeds = seeds
        |> String.split(":")
        |> List.last()
        |> String.split()
        |> Enum.map(& String.to_integer(&1))
        |> expandSeeds([])
        |> IO.inspect()

        mapper = lines
        |> getMaps(%{}, "")
        |> IO.inspect()

        stages = ["seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"]

        res = mapSeeds(mapper, stages, seeds)

        Enum.min(res)
        |> IO.inspect(label: "min")

        Enum.count(res) |> IO.inspect()

    end




    def mapSeeds(_, [], src) do src end
    def mapSeeds(mapper, [stg | tail], src) do
        m = Map.fetch!(mapper, stg) # array of tuples {left, right, func}

        x = src
        |> Enum.map(fn s->
            res = mS(s, m, [])
            |> List.flatten()
            # |> IO.inspect(label: "1 seed mapped in #{stg}")

            if res == [] do
                s
            else
                res
            end
        end)
        |> List.flatten()
        |> IO.inspect(label: "all seeds mapped in #{stg}")

        Enum.count(x) |> IO.inspect()
        mapSeeds(mapper, tail, x)
    end

    def mS(seed, map, dest) do
        map
        |> Enum.map( fn {left, right, diff} ->
            intersect(seed, {left, right})
            # |> IO.inspect(label: "intersect")
            |> groupMapper(diff, [])
        end)
    end

    def groupMapper([], _, res), do: res
    def groupMapper([{:i, l, r}| t], diff, res) do
        groupMapper(t, diff, res ++ [{l+diff, r+diff}])
    end
    def groupMapper([{:c, l, r}| t], diff, res) do
        groupMapper(t, diff, res ++ [{l, r}])
    end





    def expandSeeds([], s) do s end
    def expandSeeds([range_start | tail], s) do
        [range_len | tail] = tail
        expandSeeds(tail, s ++ [{range_start, range_start+range_len}])
    end

    def getMaps([], m, _) do m end
    def getMaps(["" | rest], m, _) do getMaps(rest, m, "") end
    def getMaps([head | rest], m, mk) when binary_part(head, 0,1) in ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"] do
        {drs, srs, r} = head
        |> String.split()
        |> Enum.map(& String.to_integer(&1))
        |> List.to_tuple() # -> {drs, srs, r}

        func = {srs, srs+r-1, drs-srs}

        m = Map.replace(m, mk, Map.fetch!(m, mk) ++ [func])
        # IO.inspect({head, m, mk}, label: "starts num")
        getMaps(rest, m, mk)
    end
    def getMaps([head | rest], m, _) do
        mk = head
        |> String.split()
        |> List.first()
        m = Map.put(m, mk, [])
        getMaps(rest, m, mk)
    end

    def intersect({s1, e1}, {s2, e2}) when s2 > e1 or s1 > e2 do [] end
    def intersect({s1, e1}, {s2, e2}) do
        inter = {:i, Enum.max([s1, s2]), Enum.min([e1,e2])}
        [inter]++[remainder({s1, e1}, inter)]
        |> List.flatten()
        |> Enum.filter(& &1!=nil)
    end

    def remainder({s1, e1}, {_, s2, e2}) when s2 > e1 or s1 > e2 do nil end
    def remainder({s1, e1}, {_, si, ei}) do
        [validateInt(Enum.min([s1, si]), Enum.max([s1, si])-1), validateInt(Enum.min([e1,ei])+1, Enum.max([e1,ei]))]
    end

    def validateInt(s, e) when s >= e do nil end
    def validateInt(s, e) do {:c, s, e} end
end

D5.solver()
