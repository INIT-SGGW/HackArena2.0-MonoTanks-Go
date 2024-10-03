Let's assume we have these jsons:
```
{
  "type": 17,
  "payload": {
      "direction": 1
  }
}
```

```
{
  "type": 18,
  "payload": {
      "rotation": 0
  }
}
```

Is there a nice way in golang to deserialize them? I mean, they come, and I dont know which one come and I would like to use one deserialization into nice structs without manually checking for fields.

Maybe I give you a pseudocode what I want to achieve:
```pseudocode
enum Message {
    Direction {
        direction: int
    }
    Rotation {
        rotation: int
    }
}

// ...

message = nextMessage() // I get message from the stream from somewhere. This is not important

parsed = parse(message)

switch (parsed) {
    if Direction {
        // now `parsed` is strictly typed as Direction class
        // This `processDirection` function requires argument of type Direction
        processDirection(parsed)
    }
    if Rotation {
        // now `parsed` is strictly typed as Rotation class
        // This `processRotation` function requires argument of type Rotation
        processRotation(parsed)
    }
}

```

