import * as fs from 'fs';

const CellType = {
    EMPTY: 0,
    SPLITTER: 1,
    ENTRYPOINT: 2
} as const;
type CellType = typeof CellType[keyof typeof CellType];

const ActiveCellType = {
    EMPTY: 0,
    INACIVESPLITTER: 1,
    ACTIVESPLITTER: 2,
    BEAM: 3
} as const;
type ActiveCellType = typeof ActiveCellType[keyof typeof ActiveCellType];

class Grid {
    private grid: CellType[][];

    constructor(fileContents: string) {
        const lines = fileContents.split('\n').filter(line => line.length > 0);
        this.grid = lines.map(line => line.split('').map(char => {
            switch (char) {
                case '.':
                    return CellType.EMPTY;
                case '^':
                    return CellType.SPLITTER;
                case 'S':
                    return CellType.ENTRYPOINT;
                default:
                    return CellType.EMPTY;
            }
        }));
    }
    
    printGrid(): void {
        for (const row of this.grid) {
            console.log(row.map(cell => {
                switch (cell) {
                    case CellType.EMPTY:
                        return '.';
                    case CellType.SPLITTER:
                        return '^';
                    case CellType.ENTRYPOINT:
                        return 'S';
                }
            }).join(''));
        }
    }
}

class ActiveGrid {
    private grid: ActiveCellType[][];
    private waysToGetHere: number[][];
    private splitCount: number = 0;

    constructor(initialGrid: Grid) {
        // Initialize active grid based on initial grid
        this.grid = []; // Placeholder initialization
        this.waysToGetHere = [];

        let firstRow = initialGrid['grid'][0].map(cell => {
            switch (cell) {
                case CellType.ENTRYPOINT:
                    return ActiveCellType.BEAM;
                default:
                    return ActiveCellType.EMPTY;
            }
        });
        this.grid.push(firstRow);
        let firstWayRow = initialGrid['grid'][0].map(cell => {
            switch (cell) {
                case CellType.ENTRYPOINT:
                    return 1;
                default:
                    return 0;
            }
        });
        this.waysToGetHere.push(firstWayRow);

        for (let i = 1; i < initialGrid['grid'].length; i++) {

            // The contents of this row are now based on the row above, and the
            // contents of the initial grid at this row.
            let newRow: ActiveCellType[] = [];
            let newWaysRow: number[] = this.waysToGetHere[i-1].slice();
            for (let j = 0; j < initialGrid['grid'][i].length; j++) {
                let aboveCell = this.grid[i - 1][j];
                let initialCell = initialGrid['grid'][i][j];
                
                if (initialCell === CellType.SPLITTER) {
                    if (aboveCell === ActiveCellType.BEAM) {
                        newRow.push(ActiveCellType.ACTIVESPLITTER);
                        this.splitCount++;
                    } else {
                        newRow.push(ActiveCellType.INACIVESPLITTER);
                    }
                } else {
                    if (aboveCell === ActiveCellType.BEAM) {
                        newRow.push(ActiveCellType.BEAM);
                    } else {
                        newRow.push(ActiveCellType.EMPTY);
                    }
                }
            }

            // After constructing the new row, we need to handle beam splitting
            for (let j = 0; j < newRow.length; j++) {
                if (newRow[j] === ActiveCellType.ACTIVESPLITTER) {
                    // Split beam to left
                    if (j > 0) {
                        newRow[j - 1] = ActiveCellType.BEAM;
                        // Use += because there might be multiple ways to get here
                        newWaysRow[j - 1] += newWaysRow[j];
                    }
                    // Split beam to right
                    if (j < newRow.length - 1) {
                        newRow[j + 1] = ActiveCellType.BEAM;
                        newWaysRow[j + 1] += newWaysRow[j];
                    }
                }
            }

            for (let j = 0; j < newRow.length; j++) {
                // If a cell is a splitter, set ways to 0
                if (newRow[j] === ActiveCellType.ACTIVESPLITTER || newRow[j] === ActiveCellType.INACIVESPLITTER) {
                    newWaysRow[j] = 0;
                }
            }
            this.grid.push(newRow);
            this.waysToGetHere.push(newWaysRow);
        }
    }

    printActiveGrid(): void {
        for (const row of this.grid) {
            console.log(row.map(cell => {
                switch (cell) {
                    case ActiveCellType.EMPTY:
                        return '.';
                    case ActiveCellType.INACIVESPLITTER:
                        return '^';
                    case ActiveCellType.ACTIVESPLITTER:
                        return '*';
                    case ActiveCellType.BEAM:
                        return '|';
                }
            }).join(''));
        }
    }

    printWaysGrid(): void {
        for (const row of this.waysToGetHere) {
            console.log(row.map(ways => ways.toString()).join(' '));
        }
    }

    countActiveSplitters(): number {
        return this.splitCount;
    }

    countActivePaths(): number {
        return this.waysToGetHere[this.waysToGetHere.length - 1].reduce((sum, ways) => sum + ways, 0);
    }
}

export function loadFileContents(filePath: string): string {
    return fs.readFileSync(filePath, 'utf-8');
}

export function runDay7Part1(fileContents: string): number {
    const initialGrid = new Grid(fileContents);
    initialGrid.printGrid();
    const activeGrid = new ActiveGrid(initialGrid);
    activeGrid.printActiveGrid();
    return activeGrid.countActiveSplitters();
}

export function runDay7Part2(fileContents: string): number {
    const initialGrid = new Grid(fileContents);
    initialGrid.printGrid();
    const activeGrid = new ActiveGrid(initialGrid);
    activeGrid.printWaysGrid();
    return activeGrid.countActivePaths();
}