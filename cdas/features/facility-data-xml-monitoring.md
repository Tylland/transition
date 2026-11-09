# Monitoring Facility Data XML Files

> As a segment service operator
> I want the service to monitor a folder for facility-data XML files  
> So that new or updated network data is processed automatically  

The service scans its configured input directory at startup and every five seconds. It considers non-directory files with the `.xml` extension and processes files that are new or have a later modification time than the last time they were seen.

## Scenario: Process XML files already present at startup
```gherkin
Given the input folder contains facility-data XML files
When the segment service starts
Then it processes the XML files found in the initial scan
```

## Scenario: Process a newly added XML file
```gherkin
Given the segment service is watching the input folder
When a new facility-data XML file is added
Then the service processes it during a subsequent scan
```

## Scenario: Reprocess a modified XML file
```gherkin
Given the service has processed a facility-data XML file
When the file's modification time becomes later than the last time it was seen
Then the service processes the file again
```

## Scenario: Ignore non-XML files and subdirectories
```gherkin
Given the segment service is scanning its input folder
When it encounters a non-XML file or a subdirectory
Then it does not process that entry as facility data
```

## Scenario: Log an unreadable or invalid XML file
```gherkin
Given the input folder contains an unreadable or invalid facility-data XML file
When the service processes that file
Then it logs the read or parse error
  And continues scanning other entries
```
