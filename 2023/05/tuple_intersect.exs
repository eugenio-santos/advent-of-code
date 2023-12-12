defmodule Tplinter do
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

Tplinter.intersect({5,5}, {5,5}) |> IO.inspect()
Tplinter.intersect({0,5}, {10,15}) |> IO.inspect()
Tplinter.intersect({5,15}, {0,10}) |> IO.inspect()
Tplinter.intersect({5,15}, {10,20}) |> IO.inspect()
Tplinter.intersect({5,15}, {0,20}) |> IO.inspect()
Tplinter.intersect({0,20}, {5,15}) |> IO.inspect()
Tplinter.intersect({0,20}, {20,30}) |> IO.inspect()
