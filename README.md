# LimaGoggles
A lightweight golang app that can take a new line separated string of text and put it into the JSON Format expected by the LimaCharlie lookup GUI.

## Usage
To use the app, run it from the command line, providing the path to the text file and the metadata string as arguments.

```
./LimaGoggles <path-to-text-file> <metadata-string>
```
## Example
This will process sample.txt, associating each line in the file with "example-metadata" and output the result in JSON format.

```
./LimaGoggles ./sample.txt "example-metadata"
```

## Output Format

The output will be in the following JSON format:
```
{
  "lineValue0": {
    "metadata": "metadata-string"
  },
  "lineValue1": {
    "metadata": "metadata-string"
  },
  // ...
}
```
