## Using Text Service

1. To 847-737-7947 text "Tell me a joke" (punctutation doesn't matter, no space or period at the end)
2. You will receive a setup (this may take up to 30 seconds)
3. Respond with "?" to receive the punchline
4. You will receive the punchline

Note: alternatively you can say "Tell me a joke about X" where X is a category, e.g. "Tell me a joke about math"


## Using Joke API
This is the backend of the text service

Public Routes:
- https://jokes.alordi.com/jokes --- Get all jokes
- https://jokes.alordi.com/jokes/type={types} --- Get jokes by type, where {types} is a list of types separated by a comma
  - ex: https://jokes.alordi.com/jokes/type=math gets all math jokes
  - ex: https://jokes.alordi.com/jokes/type=math,science gets all math and science jokes
- https://jokes.alordi.com/jokes/type!={types} --- Get jokes not in types, where {types} is a list of types separated by a comma
  - ex: https://jokes.alordi.com/jokes/type!=music gets all non-music jokes
  - ex: https://jokes.alordi.com/jokes/type=music,sports gets all jokes not in music and not in sports
- https://jokes.alordi.com/jokes/random --- Get a random joke
- https://jokes.alordi.com/jokes/random/type={types} --- Get a random joke by type, where {types} is a list of types separated by a comma
  - ex: https://jokes.alordi.com/jokes/random/type=math gets random math joke
  - ex: https://jokes.alordi.com/jokes/random/type=math,science gets random math joke or random science joke
- https://jokes.alordi.com/jokes/random/type={types} --- Get a random joke not in types, where {types} is a list of types separated by a comma
  - ex: https://jokes.alordi.com/jokes/random/type!=music gets random joke not in music
  - ex: https://jokes.alordi.com/jokes/random/type!=music,sports gets random joke not in music and not in sports
- https://jokes.alordi.com/jokes/{id} --- Get joke by id


## Joke types:
- sports
- science
- math
- music
- animals
- programming
- food
- other
