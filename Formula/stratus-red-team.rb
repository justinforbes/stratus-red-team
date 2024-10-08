# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class StratusRedTeam < Formula
  desc ""
  homepage "https://stratus-red-team.cloud"
  version "2.17.0"
  license "Apache-2.0"

  on_macos do
    on_intel do
      url "https://github.com/DataDog/stratus-red-team/releases/download/v2.17.0/stratus-red-team_Darwin_x86_64.tar.gz"
      sha256 "f24458d03e8a6d205fa8192a5d86ae298ebbb52c9dd3e12fa00828b6c5b52018"

      def install
        bin.install "stratus"

        # Install shell completions
        generate_completions_from_executable(bin/"stratus", "completion", shells: [:bash, :fish, :zsh], base_name: "stratus")
      end
    end
    on_arm do
      url "https://github.com/DataDog/stratus-red-team/releases/download/v2.17.0/stratus-red-team_Darwin_arm64.tar.gz"
      sha256 "4c9fa3ade611c76ebc551b305a67562b501bcbbbd28a91d3797a6b3587e5fbc1"

      def install
        bin.install "stratus"

        # Install shell completions
        generate_completions_from_executable(bin/"stratus", "completion", shells: [:bash, :fish, :zsh], base_name: "stratus")
      end
    end
  end

  on_linux do
    on_intel do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/DataDog/stratus-red-team/releases/download/v2.17.0/stratus-red-team_Linux_x86_64.tar.gz"
        sha256 "cc34035ac11e263bc747d3ff84c5231a6fa6ae30b99bb17d670722f35fa96bf4"

        def install
          bin.install "stratus"

          # Install shell completions
          generate_completions_from_executable(bin/"stratus", "completion", shells: [:bash, :fish, :zsh], base_name: "stratus")
        end
      end
    end
    on_arm do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/DataDog/stratus-red-team/releases/download/v2.17.0/stratus-red-team_Linux_arm64.tar.gz"
        sha256 "da87085cd04959c260e0eb0eb38de55bc572501a2cf09f0b166a9d031afdeb66"

        def install
          bin.install "stratus"

          # Install shell completions
          generate_completions_from_executable(bin/"stratus", "completion", shells: [:bash, :fish, :zsh], base_name: "stratus")
        end
      end
    end
  end
end
