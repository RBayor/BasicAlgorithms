#!/bin/bash



run_sorting_files() {
  
  local sorting_dir="./src/Sorting/bubble" # dir set to only bubble sort for now

  local js_files=$(find "$sorting_dir" -type f -name "*.js") # Find all JS files
  local go_files=$(find "$sorting_dir" -type f -name "*.go") # Find all Go files

  
  local output_dir="./dist/sorting" # Output directory

  build_go_files

  # Run the commands in parallel using background processes
  # run_with_engine "Node.js" "node" "$js_files" "$output_dir" &
  # run_with_engine "Bun.sh" "bun" "$js_files" "$output_dir" &
  # run_with_engine "Golang" "" "$go_files" "$output_dir" 

  # Wait for all background processes to finish
  # wait
}

run_with_engine() {
  local engine_name="$1"
  local engine_command="$2"
  local files="$3"
  local output_dir="$4"



  if [ "$engine_name" == "Node.js" ] || [ "$engine_name" == "Bun.sh" ]; then
    for file in $files; do
      echo "Running $file with $engine_name..."
      
      # Output of the running command
      output=$("$engine_command" "$file" 2>&1)
      
      # Display the output on the console
      echo "$output"
      
      # Save the output to a file (append mode)
      output_filename="${file##*/}.output.txt"
      output_filepath="$output_dir/$output_filename"
      mkdir -p "$output_dir"
      
      echo -e "\n\n$output" >> "$output_filepath" # Append to the file
      echo "Output appended to: $output_filepath"
      
      echo
    done
  else
    for file in files; do
      echo "building executable for $file with $engine_name..."
      go build -o "$output_dir/${file%.go}" "$file"
    done
  fi


}

build_go_files(){

  local go_files=$(find ./src/Sorting/bubble -type f -name "*.go") # Find all Go files

  for file in $go_files; do
    echo "building executable for $file with Golang..."
    output_name="${file%.go}"
    go build -o "$output_name" "$file"
  done
}




run_sorting_files

