
defmodule Solver do
    def get1stD("") do "" end
    def get1stD(s) when <<49>> <= binary_part(s, 0, 1) and binary_part(s, 0, 1) <= <<58>> do
        IO.puts(String.first(s))
    end
    def get1stD("one"<>_) do "1" end
    def get1stD("two"<>_) do "2" end
    def get1stD("three"<>_) do "3" end
    def get1stD("four"<>_) do "4" end
    def get1stD("five"<>_) do "5" end
    def get1stD("six"<>_) do "6" end
    def get1stD("seven"<>_) do "7" end
    def get1stD("eight"<>_) do "8" end
    def get1stD("nine"<>_) do "9" end
    def get1stD(s) do get1stD(String.slice(s, 1..-1)) end

    def getLstD("") do "" end
    def getLstD(s) when <<49>> <= binary_part(s, 0, 1) and binary_part(s, 0, 1) <= <<58>> do
        IO.puts(String.first(s))
    end
    def getLstD("eno"<>_) do "1" end
    def getLstD("owt"<>_) do "2" end
    def getLstD("eerht"<>_) do "3" end
    def getLstD("ruof"<>_) do "4" end
    def getLstD("evif"<>_) do "5" end
    def getLstD("xis"<>_) do "6" end
    def getLstD("neves"<>_) do "7" end
    def getLstD("thgie"<>_) do "8" end
    def getLstD("enin"<>_) do "9" end
    def getLstD(s) do getLstD(String.slice(s, 1..-1)) end

    def s do
        File.stream!("test")
        |> Stream.map(&String.trim/1)
        |> Stream.map(fn line ->
            IO.puts(line)
            # f = get1stD(line)
            IO.puts(get1stD(line))

            # IO.puts(line|> String.reverse())
            # l = getLstD(line|> String.reverse())
            # IO.puts(l)

            # IO.puts(f<>l)

            # f <> l |> String.to_integer()
        end)
        # |> Enum.reduce(fn x, acc ->
        #     IO.puts(x)
        #     IO.puts(acc)
        #     x + acc end)
        # |> IO.puts()
        |> Stream.run()
    end
end

Solver.s()
