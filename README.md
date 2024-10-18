# OSCAL and CSV

The Open Security Controls Assessment Language is a data format developed by the National Institute for Standards and Technology. It provides a machine-readable format for expressing control information. 

This machine-readable format exists to be optimized for... well machines. So how do we make it more human-friendly?

A graphical user interface would certainly be one such approach. Doing so is valiant but not a trivial undertaking. What other tooling might we have of use?

## Try it Out

### Prerequisites
- TODO: make target for build + release process

### Create

- `oac generate [MODEL]` to generate a blank CSV template

### Convert
- If you have an OSCAL model that is supported by `oac` - run `oac convert [MODEL]` to convert an existing model to CSV format

## Comma-separated Values - CSV

This repository is an experiment to see if there is any value in attempting (with data loss) to provide a structured template in CSV format that could be easily translated to OSCAL. 

This would enable tooling that can graphically visualize CSV data to present control information for modification by humans manually with the intent that changes are then reconciled back to OSCAL format. 

## Data Loss

OSCAL provides many nested or otherwise structured models, objects, and fields that are hard to capture in CSV format. Initial attempts at this project will be under the intent of translating the "core" information from a model to and from CSV format. 

What we can do to improve the potential for data-loss is performing a merge when the data exists in an existing OSCAL model. 

See the [docs](./docs/README.md) for more information on model specifics.

## Future

Implement a system to support nested fields. Thinking along the lines of a single CSV field that creates a delimited string in such a way as to define nested field key/value.