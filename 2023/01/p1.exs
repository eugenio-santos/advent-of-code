defmodule Solver do
  def s do
    File.stream!("puzzle")
    |> Stream.map(fn line ->
      nums =
        line
        |> String.replace(~r/[^\d]/, "")
        |> String.trim()

      (String.first(nums) <> String.last(nums))
      |> String.to_integer()
    end)
    |> Enum.sum()
    |> IO.puts()
  end
end

Solver.s()
