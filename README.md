```
 _____       _            _           _                 
/  __ \     | |          | |         | |                
| /  \/ __ _| | ___ _   _| | __ _  __| | ___  _ __ __ _ 
| |    / _` | |/ __| | | | |/ _` |/ _` |/ _ \| '__/ _` |
| \__/\ (_| | | (__| |_| | | (_| | (_| | (_) | | | (_| |
 \____/\__,_|_|\___|\__,_|_|\__,_|\__,_|\___/|_|  \__,_|
                                                        
                                                        
 _____                          _                       
|_   _|                        | |                      
  | | _ __ ___  _ __   ___  ___| |_ ___                 
  | || '_ ` _ \| '_ \ / _ \/ __| __/ _ \                
 _| || | | | | | |_) | (_) \__ \ || (_) |               
 \___/_| |_| |_| .__/ \___/|___/\__\___/                
               | |                                      
               |_|                                      
______                                                  
|  _  \                                                 
| | | |_____   __                                       
| | | / _ \ \ / /                                       
| |/ /  __/\ V /                                        
|___/ \___| \_/                                         
                                                        
                                                        
______                     _                            
| ___ \                   | |                           
| |_/ /___ _ __ ___   ___ | |_ ___                      
|    // _ \ '_ ` _ \ / _ \| __/ _ \                     
| |\ \  __/ | | | | | (_) | || (_) |                    
\_| \_\___|_| |_| |_|\___/ \__\___/                     
                                                      
```

# Sobre 

* Calculadora de imposto para devs que trabalham no Brasil e que emitem nota para clientes no exterior

* Inspirada no artigo https://medium.com/@hnordt/como-pagar-6-de-impostos-dentro-da-lei-a91c23868ec6

# Utilização

* Demo em funcionamento:
    * `https://ooh86uy8d7.execute-api.us-east-1.amazonaws.com/dev?faturamento-anual=120000`
    (altere o parâmetro `faturamento-anual` da URL)

* A saída JSON possui 3 blocos principais: "input", "output" e "sobre"
    * O bloco `output` apresenta o resultado da simulação

* Dica: utilize um formatador de JSON para a saída ficar mais amigável (exemplo: https://jsonformatter.org/)

# Instalação customizada

* Para você modificar e rodar a sua própria instância dessa função serverless:
    * Crie uma conta na AWS
    * Configure o aws-cli com as suas chaves de acesso
    * Clone este repositório
    * Instale o framework "serverless" (https://www.serverless.com/)
    * Copie o arquivo env.json.dist para env.json e altere o AWS_PROFILE se necessário
    * Rode o comando `make deploy` (ou então `make build` + `serverless deploy`)
    * O cli irá exibir o endpoint da sua lambda function