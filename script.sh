#!/bin/bash

run_sorting_files() {
  
  local sorting_dir="./src/Sorting/bubble" # dir set to only bubble sort for now

  local js_files=$(find "$sorting_dir" -type f -name "*.js") # Find all JS files
  local go_files=$(find "$sorting_dir" -type f -name "*.go" -exec realpath {} \;) #Find all Go files
  local go_files=$(find ./src/Sorting/bubble -type f -name "*.go") # Find all Go files in bubble sort dir for now

  # TODO: Add a function to execute data-gen.js to generate data.json files
  build_go_files "$go_files" &
  wait
  
  run_with_engine "Node.js" "node" "$js_files"
  run_with_engine "Bun.sh" "bun" "$js_files"
  run_with_engine "Golang" "go" "$go_files" 

}

run_with_engine() {
  local engine_name="$1"
  local engine_command="$2"
  local files="$3"
  local version=""

  if [ "$engine_name" == "Node.js" ]; then
    version="engine_name: $engine_name, version: $("$engine_command" --version) 🔵"
  elif [ "$engine_name" == "Bun.sh" ]; then
    version="engine_name: $engine_name, version: $("$engine_command" --version) ⚪️"
  elif [ "$engine_name" == "Golang" ]; then
    version="engine_name: $engine_name, version: $("$engine_command" version) 🟢"
  fi

  if [ "$engine_name" == "Node.js" ] || [ "$engine_name" == "Bun.sh" ]; then
    run_with_node_or_bun "$engine_name" "$engine_command" "$files" "$version"
  else
    run_go_executable "$engine_name" "$files" "$version"
  fi
}

run_with_node_or_bun() {
  local engine_name="$1"
  local engine_command="$2"
  local files="$3"
  local version="$4"

  for file in $files; do
    echo
    echo "Running $file with $version..."
    echo
    
    base_name="$(basename "$file" .js)"
    dir_path="${file%/*}" 

    # Output of the running command
    # output=$("$engine_command" "$file" 2>&1)
    output=$(echo -e "$version\n$("$engine_command" "$file" 2>&1)")

    # Display the output on the console
    echo "$output"
    echo

    # Save the output to a file (append mode)
    output_filename="${base_name}.output.txt"
    output_filepath="$dir_path/$output_filename"
    mkdir -p "$dir_path"

    echo -e "\n\n$output" >> "$output_filepath" # Append to the file
    echo "Output appended to: $output_filepath"
  done
}

run_go_executable() {
  local engine_name="$1"
  local files="$2"
  local version="$3"

  for file in $files; do
    base_name="$(basename "$file" .go)"
    dir_path="${file%/*}" # Get the directory path

    data_path=$(find "$dir_path" -type f -name "*.json" -print -quit)
    exec_file=$(find "$dir_path" -type f -perm +111 -print -quit)

    echo
    echo "Running $exec_file with $engine_name... $version"
    echo

    # Check if an executable and data were found
    if [ -n "$exec_file" ] && [ -n "$data_path" ]; then
      data_path=$(realpath "$data_path") # Get the absolute path

      # output=$("$exec_file" "$data_path" 2>&1) # Run the executable
       output=$(echo -e "$version\n$("$exec_file" "$data_path" 2>&1)")

      # Display the output on the console
      echo "$output"
      echo

      # Save the output to a file (append mode)
      output_filename="${base_name}.output.txt"
      output_filepath="$dir_path/$output_filename"
      mkdir -p "$dir_path"

      echo -e "\n\n$output" >> "$output_filepath" # Append to the file
      echo "Output appended to: $output_filepath"
    else
      echo "Executable or data missing in $dir_path"
    fi
  done
}


build_go_files(){
  go_files="$1"
  
  for file in $go_files; do
    base_name="$(basename "$file" .js)"
    echo "⚙️ Building go binary for $base_name"
    output_name="${file%.go}"
    go build -o "$output_name" "$file"
  done
  echo "✅ Done building go binary"
}


run_sorting_files

# run_with_engine() {
#   local engine_name="$1"
#   local engine_command="$2"
#   local files="$3"
#   local output_dir="$4"

#   if [ "$engine_name" == "Node.js" ] || [ "$engine_name" == "Bun.sh" ]; then
#     for file in $files; do
#       echo
#       echo "Running $file with $engine_name..."
#       echo
#       base_name="$(basename "$file" .js)" 
      
#       # Output of the running command
#       output=$("$engine_command" "$file" 2>&1)
      
#       # Display the output on the console
#       echo "$output"
#       echo
      
#       # Save the output to a file (append mode)
#       output_filename="${base_name}.output.txt"
#       output_filepath="$output_dir/$output_filename"
#       mkdir -p "$output_dir"
      
#       echo -e "\n\n$output" >> "$output_filepath" # Append to the file
#       echo "Output appended to: $output_filepath"
#       echo
#     done
#   else
#     for file in $files; do
#         base_name="$(basename "$file" .go)"         
#         dir_path="${file%/*}" # Get the directory path
        
#         data_path=$(find "$dir_path" -type f -name "*.json" -print -quit)
#         exec_file=$(find "$dir_path" -type f -perm +111 -print -quit)

#         echo
#         echo "Running $exec_file with $engine_name..."
#         echo

#         # Check if an executable and data was found
#         if [ -n "$exec_file" ] && [ -n "$data_path" ]; then
            
#             data_path=$(realpath "$data_path") # Get the absolute path

#             output=$("$exec_file" "$data_path" 2>&1) # Run the executable

#             # Display the output on the console
#             echo "$output"
#             echo

#             # Save the output to a file (append mode)
#             output_filename="${base_name}.output.txt"
#             output_filepath="$output_dir/$output_filename"
#             mkdir -p "$output_dir"

#             echo -e "\n\n$output" >> "$output_filepath" # Append to the file
#             echo "Output appended to: $output_filepath"
#         else
#             echo "Executable or data missin in $dir_path"
#         fi
#     done
#   fi
# }
