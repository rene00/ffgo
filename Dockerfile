FROM swaggerapi/swagger-codegen-cli-v3:3.0.30

CMD ["generate", "-i", "https://api-docs.firefly-iii.org/firefly-iii-1.5.4.yaml", "-l", "go", "-o", "/local/out/go", "--additional-properties", "packageName=ffgo"]
