class Fcat < Formula
  desc "Ferramenta CLI que exibe citações de Nietzsche em formato ASCII."
  homepage "https://github.com/seu_usuario/seu_repositorio"  # Atualize com seu link
  url "https://github.com/seu_usuario/seu_repositorio/releases/download/v0.1.0/fcat"  # Atualize com o link para o executável
  version "0.1.0"
  sha256 "SHA256_DO_EXECUTÁVEL"  # Você pode usar o comando `shasum -a 256 fcat` para gerar

  def install
    bin.install "fcat"
  end

  test do
    system "#{bin}/fcat", "--version"
  end
end

