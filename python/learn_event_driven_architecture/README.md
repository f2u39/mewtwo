# Learn Event-Driven Architecture

Gonna learn [Event-Driven Architecture with React and FastAPI – Full Course](https://www.youtube.com/watch?v=NVvIpqmf_Xc&t=1540s&ab_channel=freeCodeCamp.org)

## Course code

- [Front](https://github.com/scalablescripts/react-event-driven)
- [Back](https://github.com/scalablescripts/fast-api-event-driven)

## Env

``` cmd
brew install pyenv

pyenv install --list

pyenv install 3.10.6

pyenv versions

pyenv global 3.10.6

vim ~/.bash_profile
    eval "$(pyenv init --path)"

source ~/.bash_profile
```

## Run

``` cmd
pip install uvicorn

uvicorn main:app --reload
```

## Try

``` cmd
Mehtod: POST
URL: http://localhost:8000/deliveries/create
Body: 
{
    "type": "CREATE_DELIVERY",
    "data": {
        "budget": 150000,
        "notes": "iPhone 14 Pro"
    }
}
```
