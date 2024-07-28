# Catalog

The catalog translation to/from CSV to OSCAL is "Control Centric". This means that it focuses on the generation of controls in either isolation or within a group. 

## Expectations

This conversion will focus on conversion of `controls` and `groups` from OSCAL to CSV rows.

`Group` column will be an optional identifier that groups controls together into a `group`. All controls without a `group` column entry will be considered individual controls. 