defmodule D18.MixProject do
    use Mix.Project

    def project do
      [
        app: :d18,
        version: "0.0.1",
        deps: deps(),
      ]
    end

    def deps do
        [
          {:matrex, "~> 0.6"}
        ]
    end
end
