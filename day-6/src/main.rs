use std::fs::File;
use std::io::{BufRead, BufReader};

enum Operation {
    Add,
    Multiply,
}

struct CalculationInput {
    operation: Operation,
    values: Vec<i64>,
}

impl CalculationInput {
    fn default() -> Self {
        CalculationInput {
            operation: Operation::Add,
            values: Vec::new(),
        }
    }

    fn default_with_op(op: Operation) -> Self {
        CalculationInput {
            operation: op,
            values: Vec::new(),
        }
    }

    fn evaluate(&self) -> i64 {
        match self.operation {
            Operation::Add => self.values.iter().sum(),
            Operation::Multiply => self.values.iter().product(),
        }
    }
}

fn setup_operations(lines: &mut Vec<String>) -> Result<Vec<CalculationInput>, &'static str> {
    // Create the calculations array from the operations line
    Ok(lines.pop().ok_or("Could not extract operators line")?.split_whitespace().map(
        |op_char| {
            match op_char {
                "*" => CalculationInput::default_with_op(Operation::Multiply),
                "+" => CalculationInput::default_with_op(Operation::Add),
                _ => CalculationInput::default(),
            }
        }
    ).collect())
}

trait CephalopodMathsConverter {
    fn convert_to_calculations(&self, lines: Vec<String>) -> Result<Vec<CalculationInput>, &'static str>;
}

#[allow(unused)]
struct Part1CephalopodMathsConverter;
impl CephalopodMathsConverter for Part1CephalopodMathsConverter {
    fn convert_to_calculations(&self, mut lines: Vec<String>) -> Result<Vec<CalculationInput>, &'static str> {
        let mut calculations = setup_operations(&mut lines)?;

        lines.iter().for_each(|line| {
            let mut calc_index = 0;
            line.split_whitespace().for_each(|part| {
                let value: i64 = part.parse().unwrap();
                calculations[calc_index].values.push(value);
                calc_index += 1;
            });
        });

        Ok(calculations)
    }
}

#[allow(unused)]
struct Part2CephalopodMathsConverter;
impl CephalopodMathsConverter for Part2CephalopodMathsConverter {
    fn convert_to_calculations(&self, mut lines: Vec<String>) -> Result<Vec<CalculationInput>, &'static str> {
        let mut calculations = setup_operations(&mut lines)?;

        let line_length = lines.iter().map(|line| line.len()).max().unwrap_or(0);
        let mut calc_index = 0;

        for i in 0..line_length {
            let mut values_in_column = String::new();
            lines.iter().for_each(|line| {
                if let Some(ch) = line.chars().nth(i) && ch != ' ' {
                    values_in_column.push(ch);
                }
            });
            if values_in_column.is_empty() {
                values_in_column.clear();
                calc_index += 1;
            }
            else {
                let value: i64 = values_in_column.parse().unwrap();
                calculations[calc_index].values.push(value);
            }
        }
        Ok(calculations)
    }
}

fn read_file(filename: String) -> Result<Vec<String>, &'static str> {

    let file = File::open(filename);
    if let Err(_) = file {
        Err("Could not open file")
    } else {
        let file = file.unwrap();
        let reader = BufReader::new(file);
        reader.lines().collect::<Result<Vec<_>,_>>().map_err(|_| "Could not read lines from file")
    }
}

fn main() {

    let file = std::env::args().nth(1).expect("No file specified");
    let converter = Part2CephalopodMathsConverter;

    let lines = read_file(file);
    if lines.is_err() {
        eprintln!("Error: {}", lines.err().unwrap());
        return;
    }

    match converter.convert_to_calculations(lines.unwrap()) {
        Ok(calcs) => {
            println!("Result: {}", calcs.iter().map(|calc| calc.evaluate()).sum::<i64>());
        },
        Err(e) => {
            eprintln!("Error: {}", e);
        }
    }

}
