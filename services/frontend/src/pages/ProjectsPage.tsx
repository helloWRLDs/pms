import { FC, useState } from "react";

const ProjectsPage: FC = () => {
  const [searchQuery, setSearchQuery] = useState("");

  const projects = [
    {
      id: 1,
      title: "Website Redesign",
      status: "Active",
      progress: 75,
      members: [
        {
          name: "John Doe",
          avatar:
            "https://readdy.ai/api/search-image?query=professional%20headshot%20of%20a%20young%20business%20person%20with%20a%20warm%20smile%20wearing%20formal%20attire%20against%20a%20neutral%20background%20perfect%20for%20corporate%20profile&width=40&height=40&seq=1&orientation=squarish",
        },
        {
          name: "Jane Smith",
          avatar:
            "https://readdy.ai/api/search-image?query=professional%20headshot%20of%20a%20female%20executive%20with%20confident%20expression%20wearing%20business%20suit%20against%20clean%20studio%20background&width=40&height=40&seq=2&orientation=squarish",
        },
        {
          name: "Mike Johnson",
          avatar:
            "https://readdy.ai/api/search-image?query=corporate%20portrait%20of%20a%20middle%20aged%20businessman%20with%20glasses%20against%20modern%20office%20background%20showing%20professionalism&width=40&height=40&seq=3&orientation=squarish",
        },
      ],
      tasksCompleted: 15,
      totalTasks: 20,
      deadline: "2025-05-01",
    },
    {
      id: 2,
      title: "Mobile App Development",
      status: "On Hold",
      progress: 45,
      members: [
        {
          name: "Sarah Wilson",
          avatar:
            "https://readdy.ai/api/search-image?query=professional%20headshot%20of%20a%20young%20female%20professional%20with%20natural%20makeup%20wearing%20modern%20business%20attire%20against%20light%20background&width=40&height=40&seq=4&orientation=squarish",
        },
        {
          name: "Tom Brown",
          avatar:
            "https://readdy.ai/api/search-image?query=corporate%20portrait%20of%20a%20young%20male%20professional%20with%20friendly%20smile%20wearing%20suit%20against%20clean%20background&width=40&height=40&seq=5&orientation=squarish",
        },
      ],
      tasksCompleted: 30,
      totalTasks: 50,
      deadline: "2025-06-15",
    },
    {
      id: 3,
      title: "Marketing Campaign",
      status: "Completed",
      progress: 100,
      members: [
        {
          name: "Emily Davis",
          avatar:
            "https://readdy.ai/api/search-image?query=professional%20business%20portrait%20of%20a%20confident%20woman%20executive%20in%20modern%20office%20wear%20against%20neutral%20backdrop&width=40&height=40&seq=6&orientation=squarish",
        },
        {
          name: "David Lee",
          avatar:
            "https://readdy.ai/api/search-image?query=corporate%20headshot%20of%20an%20asian%20businessman%20with%20professional%20appearance%20against%20clean%20studio%20setting&width=40&height=40&seq=7&orientation=squarish",
        },
        {
          name: "Lisa Anderson",
          avatar:
            "https://readdy.ai/api/search-image?query=professional%20portrait%20of%20a%20female%20business%20leader%20with%20warm%20smile%20wearing%20elegant%20outfit%20against%20light%20background&width=40&height=40&seq=8&orientation=squarish",
        },
      ],
      tasksCompleted: 40,
      totalTasks: 40,
      deadline: "2025-04-30",
    },
  ];
  const getStatusColor = (status: string) => {
    switch (status.toLowerCase()) {
      case "active":
        return "bg-[rgb(41,43,41)]";
      case "on hold":
        return "bg-[rgb(31,33,31)]";
      case "completed":
        return "bg-[rgb(26,28,26)]";
      default:
        return "bg-gray-500";
    }
  };
  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
      year: "numeric",
    });
  };
  return (
    <div className="min-h-screen bg-primary-200">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Header */}
        <div className="flex items-center justify-between mb-8">
          <div className="flex items-center">
            <i className="fas fa-project-diagram text-2xl text-[rgb(41,43,41)] mr-3"></i>
            <h1 className="text-2xl font-bold text-muted-100">Projects</h1>
          </div>
          <div className="flex items-center space-x-4">
            <div className="relative">
              <input
                type="text"
                placeholder="Search projects..."
                className="w-64 pl-10 pr-4 py-2 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[rgb(41,43,41)] focus:border-transparent text-muted-200"
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
              />
              <i className="fas fa-search absolute left-3 top-1/2 transform -translate-y-1/2 text-m-400 text-sm"></i>
            </div>
            <button className="inline-flex items-center px-4 py-2 border border-gray-300 rounded-lg text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 cursor-pointer whitespace-nowrap !rounded-button">
              <i className="fas fa-filter mr-2"></i>
              Filter
            </button>
            <button className="inline-flex items-center px-4 py-2 border border-transparent rounded-lg text-sm font-medium text-white bg-[rgb(41,43,41)] hover:bg-[rgb(31,33,31)] cursor-pointer whitespace-nowrap !rounded-button">
              <i className="fas fa-plus mr-2"></i>
              Create New Project
            </button>
          </div>
        </div>
        {/* Project Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {projects.map((project) => (
            <div
              key={project.id}
              className="bg-primary-200 rounded-lg shadow-sm hover:shadow-md transition-shadow duration-200 cursor-pointer overflow-hidden"
            >
              <div className="p-6">
                <div className="flex items-center justify-between mb-4">
                  <h3 className="text-lg font-semibold text-gray-900 truncate">
                    {project.title}
                  </h3>
                  <button className="text-gray-400 hover:text-gray-600 cursor-pointer">
                    <i className="fas fa-ellipsis-h"></i>
                  </button>
                </div>
                <div className="flex items-center mb-4">
                  <span
                    className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${getStatusColor(
                      project.status
                    )} text-white`}
                  >
                    {project.status}
                  </span>
                </div>
                <div className="mb-4">
                  <div className="flex items-center justify-between mb-2">
                    <span className="text-sm font-medium text-gray-700">
                      Progress
                    </span>
                    <span className="text-sm font-medium text-gray-700">
                      {project.progress}%
                    </span>
                  </div>
                  <div className="w-full bg-gray-200 rounded-full h-2">
                    <div
                      className="bg-[rgb(41,43,41)] rounded-full h-2 transition-all duration-300"
                      style={{ width: `${project.progress}%` }}
                    ></div>
                  </div>
                </div>
                <div className="flex items-center justify-between mb-4">
                  <div className="flex -space-x-2">
                    {project.members.slice(0, 4).map((member, index) => (
                      <img
                        key={index}
                        className="w-8 h-8 rounded-full border-2 border-white object-cover"
                        src={member.avatar}
                        alt={member.name}
                        title={member.name}
                      />
                    ))}
                    {project.members.length > 4 && (
                      <div className="w-8 h-8 rounded-full bg-gray-100 border-2 border-white flex items-center justify-center">
                        <span className="text-xs font-medium text-gray-600">
                          +{project.members.length - 4}
                        </span>
                      </div>
                    )}
                  </div>
                </div>
                <div className="flex items-center justify-between text-sm text-gray-500">
                  <div className="flex items-center">
                    <i className="fas fa-tasks mr-2"></i>
                    <span>
                      {project.tasksCompleted}/{project.totalTasks} tasks
                    </span>
                  </div>
                  <div className="flex items-center">
                    <i className="far fa-calendar mr-2"></i>
                    <span>{formatDate(project.deadline)}</span>
                  </div>
                </div>
              </div>
            </div>
          ))}
        </div>
        {/* Empty State */}
        {projects.length === 0 && (
          <div className="text-center py-12">
            <div className="mb-4">
              <i className="fas fa-folder-open text-6xl text-gray-300"></i>
            </div>
            <h3 className="text-lg font-medium text-gray-900 mb-2">
              No projects found
            </h3>
            <p className="text-gray-500 mb-6">
              Get started by creating a new project
            </p>
            <button className="inline-flex items-center px-4 py-2 border border-transparent rounded-lg text-sm font-medium text-white bg-[rgb(41,43,41)] hover:bg-[rgb(31,33,31)] cursor-pointer whitespace-nowrap !rounded-button">
              <i className="fas fa-plus mr-2"></i>
              Create New Project
            </button>
          </div>
        )}
      </div>
    </div>
  );
};

export default ProjectsPage;
