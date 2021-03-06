package genalgo

//The main structure of the library
//pSize is the size of a null population

//population[] is current population
//reproduction[] is population of descendants and mutants
//reproduction[] and population[] are separated for some selection schemas

//totalFitness is sum of fitness of both parents and descendants and mutants
//iteration is current population number

//maxIteration is max iteration that could be reached
//set maxIteration to math.Inf, if you don't need it

//Generator is chosen IGenerator implementation
//Select is chosen ISelect implementation
//Crossover is chosen ICrossover implementation
//Mutate is chosen IMutator implementation

//Fitness() is function that determines unit's fitness
//Fitness()  can be changed to fit your needs

type GenAlgo struct {
	pSize        int
	Populaion    []IUnit
	reproduction []IUnit
	totalFitness float64
	iteration    int
	MaxIteration int

	Generator IGenerator
	Select    ISelector
	Crossover ICrossover
	Mutate    IMutate
	Schema 	  ISchema

	//оператор приспособленности
	Fitness func(IUnit) float64
	//что делается в начале каждой итерации
	OnBegin func()
	//что делается в конце каждой итерации
	OnEnd func()
}

func (ga *GenAlgo) nullPopulation() {
	ga.iteration = 0
	var population []IUnit
	for i := 0; i < ga.pSize; i++ {
		p := BaseUnit{}
		cromo := ga.Generator.Generate()
		p.SetCromosomes(cromo)
		f := ga.Fitness(&p)
		p.SetFitness(f)
		ga.totalFitness += f
		population = append(population, &p)
	}
	ga.Populaion = population
}

func (ga *GenAlgo) Init(populationSize int) {
	//получить начальную
	ga.pSize = populationSize
	ga.nullPopulation()
}

func (ga *GenAlgo) ExitOn() bool {
	return ga.MaxIteration == ga.iteration
}

func (ga *GenAlgo) NextGeneration() {
	ga.iteration += 1

	var repro []IUnit
	//определить размер поколений
	N := int(float64(ga.pSize) * ga.Crossover.GetSpeed())
	for i := 0; i <= N; i += 2 {
		A, B := ga.Select.Mater(ga.Populaion)
		C, D := ga.Crossover.Cross(A, B)
		C.SetFitness(ga.Fitness(C))
		D.SetFitness(ga.Fitness(D))
		repro = append(repro, C, D)
	}

	M := int(float64(ga.pSize) * ga.Mutate.GetSpeed())
	for i := 0; i <= M; i += 2 {
		A := ga.Select.Mutator(ga.Populaion)
		B := ga.Select.Mutator(ga.Populaion)
		C := ga.Mutate.Mutate(A)
		D := ga.Mutate.Mutate(B)
		C.SetFitness(ga.Fitness(C))
		D.SetFitness(ga.Fitness(D))
		repro = append(repro, C, D)
	}

	ga.reproduction = repro
}

func (ga *GenAlgo) Simulation() {

	for !ga.ExitOn() {
		if ga.OnBegin != nil {
			ga.OnBegin()
		}
		ga.NextGeneration()
		ga.Populaion = ga.Schema.Create(ga.Populaion, ga.reproduction)
		if ga.OnEnd != nil {
			ga.OnEnd()
		}
	}
}
