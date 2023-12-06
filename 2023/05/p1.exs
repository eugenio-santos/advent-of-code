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

        mapper = lines
        |> getMaps(%{}, "")

        stages = ["seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"]

        seeds
        |> Enum.map(fn s ->
            mapSeed(mapper, stages, s)
        end)
        |> Enum.min()
        |> IO.inspect()

    end

    def mapSeed(_, [], src) do src end
    def mapSeed(mapper, [stg | tail], src) do
        m = Map.fetch!(mapper, stg) # array of tuples {drs, srs, rl}
        |> Enum.find(fn {_, srs, rl} ->
            src in srs..srs+rl-1
        end)

        if m == nil do
            mapSeed(mapper, tail, src)
        else
            diff = elem(m, 0) - elem(m, 1)
            mapSeed(mapper, tail, src+diff)
        end
    end


    def getMaps([], m, _) do m end
    def getMaps(["" | rest], m, _) do getMaps(rest, m, "") end
    def getMaps([head | rest], m, mk) when binary_part(head, 0,1) in ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"] do
        func = head
        |> String.split()
        |> Enum.map(& String.to_integer(&1))
        |> List.to_tuple()
        m = Map.replace(m, mk, Map.fetch!(m, mk) ++ [func])
        # IO.inspect({head, m, mk}, label: "starts num")
        getMaps(rest, m, mk)
    end
    def getMaps([head | rest], m, _) do
        mk = head
        |> String.split()
        |> List.first()
        m = Map.put(m, mk, [])
        # IO.inspect({head, m, mk}, label: "starts key")
        getMaps(rest, m, mk)
    end

end

D5.solver()
